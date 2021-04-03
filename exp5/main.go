package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
)

func main() {
	flag.Parse()

	// Many tools pass their command-line arguments (after any flags)
	// uninterpreted to packages.Load so that it can interpret them
	// according to the conventions of the underlying build system.
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedName}
	fmt.Println("args:", flag.Args())
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Print the names of the source files
	// for each package listed on the command line.
	for _, pkg := range pkgs {
		fmt.Println(pkg.ID, pkg.Name, pkg.GoFiles)
		fmt.Println("9999--------", pkg.PkgPath)
		g := &gen{
			pkg:         pkg,
			anonImports: make(map[string]bool),
			imports:     make(map[string]importInfo),
			values:      make(map[ast.Expr]string),
		}
		generateInjectors(g, pkg)
	}
}

// importInfo holds info about an import.
type importInfo struct {
	// name is the identifier that is used in the generated source.
	name string
	// differs is true if the import is given an identifier that does not
	// match the package's identifier.
	differs bool
}

type gen struct {
	pkg         *packages.Package
	buf         bytes.Buffer
	imports     map[string]importInfo
	anonImports map[string]bool
	values      map[ast.Expr]string
}

func newGen(pkg *packages.Package) *gen {
	return &gen{
		pkg:         pkg,
		anonImports: make(map[string]bool),
		imports:     make(map[string]importInfo),
		values:      make(map[ast.Expr]string),
	}
}

type objCacheEntry struct {
	val  interface{} // *Provider, *ProviderSet, *IfaceBinding, or *Value
	errs []error
}

type objRef struct {
	importPath string
	name       string
}

type objectCache struct {
	fset     *token.FileSet
	packages map[string]*packages.Package
	objects  map[objRef]objCacheEntry
	hasher   typeutil.Hasher
}

func newObjectCache(pkgs []*packages.Package) *objectCache {
	if len(pkgs) == 0 {
		panic("object cache must have packages to draw from")
	}
	oc := &objectCache{
		fset:     pkgs[0].Fset,
		packages: make(map[string]*packages.Package),
		objects:  make(map[objRef]objCacheEntry),
		hasher:   typeutil.MakeHasher(),
	}
	// Depth-first search of all dependencies to gather import path to
	// packages.Package mapping. go/packages guarantees that for a single
	// call to packages.Load and an import path X, there will exist only
	// one *packages.Package value with PkgPath X.
	stk := append([]*packages.Package(nil), pkgs...)
	for len(stk) > 0 {
		p := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if oc.packages[p.PkgPath] != nil {
			continue
		}
		oc.packages[p.PkgPath] = p
		for _, imp := range p.Imports {
			stk = append(stk, imp)
		}
	}
	return oc
}

// errorCollector manages a list of errors. The zero value is an empty list.
type errorCollector struct {
	errors []error
}

// add appends any non-nil errors to the collector.
func (ec *errorCollector) add(errs ...error) {
	for _, e := range errs {
		if e != nil {
			ec.errors = append(ec.errors, e)
		}
	}
}

func isWireImport(path string) bool {
	// TODO(light): This is depending on details of the current loader.
	const vendorPart = "vendor/"
	if i := strings.LastIndex(path, vendorPart); i != -1 && (i == 0 || path[i-1] == '/') {
		path = path[i+len(vendorPart):]
	}
	return path == "github.com/google/wire"
}

func qualifiedIdentObject(info *types.Info, expr ast.Expr) types.Object {
	switch expr := expr.(type) {
	case *ast.Ident:
		return info.ObjectOf(expr)
	case *ast.SelectorExpr:
		pkgName, ok := expr.X.(*ast.Ident)
		if !ok {
			return nil
		}
		if _, ok := info.ObjectOf(pkgName).(*types.PkgName); !ok {
			return nil
		}
		return info.ObjectOf(expr.Sel)
	default:
		return nil
	}
}

// findInjectorBuild returns the wire.Build call if fn is an injector template.
// It returns nil if the function is not an injector template.
func findInjectorBuild(info *types.Info, fn *ast.FuncDecl) (*ast.CallExpr, error) {
	if fn.Body == nil {
		return nil, nil
	}
	numStatements := 0
	invalid := false
	var wireBuildCall *ast.CallExpr
	for _, stmt := range fn.Body.List {
		switch stmt := stmt.(type) {
		case *ast.ExprStmt:
			numStatements++
			if numStatements > 1 {
				invalid = true
			}
			call, ok := stmt.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			if qualifiedIdentObject(info, call.Fun) == types.Universe.Lookup("panic") {
				if len(call.Args) != 1 {
					continue
				}
				call, ok = call.Args[0].(*ast.CallExpr)
				if !ok {
					continue
				}
			}
			buildObj := qualifiedIdentObject(info, call.Fun)
			if buildObj == nil || buildObj.Pkg() == nil || !isWireImport(buildObj.Pkg().Path()) || buildObj.Name() != "Build" {
				continue
			}
			wireBuildCall = call
		case *ast.EmptyStmt:
			// Do nothing.
		case *ast.ReturnStmt:
			// Allow the function to end in a return.
			if numStatements == 0 {
				return nil, nil
			}
		default:
			invalid = true
		}

	}
	if wireBuildCall == nil {
		return nil, nil
	}
	if invalid {
		return nil, errors.New("a call to wire.Build indicates that this function is an injector, but injectors must consist of only the wire.Build call and an optional return")
	}
	return wireBuildCall, nil
}

func getInterface(x ast.Decl) {
	if x, ok := x.(*ast.GenDecl); ok {
		if x.Tok != token.TYPE {
			return
		}
		for _, x := range x.Specs {
			if x, ok := x.(*ast.TypeSpec); ok {
				iname := x.Name
				if x, ok := x.Type.(*ast.InterfaceType); ok {
					for _, x := range x.Methods.List {
						if len(x.Names) == 0 {
							return
						}
						mname := x.Names[0].Name
						fmt.Println("interface:", iname, "method:", mname)

					}
				}
			}
		}
	}

}

// generateInjectors generates the injectors for a given package.
func generateInjectors(g *gen, pkg *packages.Package) (injectorFiles []*ast.File, _ []error) {
	//oc := newObjectCache([]*packages.Package{pkg})
	injectorFiles = make([]*ast.File, 0, len(pkg.Syntax))
	ec := new(errorCollector)
	fmt.Println(pkg.Syntax, "pkg id name  path:", pkg.ID, pkg.Name, pkg.PkgPath)
	for _, f := range pkg.Syntax {
		for _, decl := range f.Decls {
			fmt.Println(decl)
			getInterface(decl)

			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			fmt.Println(fn.Recv, fn.Name, fn.Type)
			buildCall, err := findInjectorBuild(pkg.TypesInfo, fn)
			if err != nil {
				ec.add(err)
				continue
			}
			if buildCall == nil {
				continue
			}
			if len(injectorFiles) == 0 || injectorFiles[len(injectorFiles)-1] != f {
				// This is the first injector generated for this file.
				// Write a file header.
				name := filepath.Base(g.pkg.Fset.File(f.Pos()).Name())
				g.p("// Injectors from %s:\n\n", name)
				injectorFiles = append(injectorFiles, f)
			}
			sig := pkg.TypesInfo.ObjectOf(fn.Name).Type().(*types.Signature)
			ins, _, err := injectorFuncSignature(sig)
			fmt.Println(sig, ins)
			if err != nil {
				/*if w, ok := err.(*wireErr); ok {
					ec.add(notePosition(w.position, fmt.Errorf("inject %s: %v", fn.Name.Name, w.error)))
				} else {
					ec.add(notePosition(g.pkg.Fset.Position(fn.Pos()), fmt.Errorf("inject %s: %v", fn.Name.Name, err)))
				}
				*/
				continue
			}
			/*injectorArgs := &InjectorArgs{
				Name:  fn.Name.Name,
				Tuple: ins,
				Pos:   fn.Pos(),
			}
			set, errs := oc.processNewSet(pkg.TypesInfo, pkg.PkgPath, buildCall, injectorArgs, "")
			if len(errs) > 0 {
				ec.add(notePositionAll(g.pkg.Fset.Position(fn.Pos()), errs)...)
				continue
			}

			if errs := g.inject(fn.Pos(), fn.Name.Name, sig, set, fn.Doc); len(errs) > 0 {
				ec.add(errs...)
				continue
			}
			*/
		}

		for _, impt := range f.Imports {
			if impt.Name != nil && impt.Name.Name == "_" {
				g.anonImports[impt.Path.Value] = true
			}
		}
	}
	if len(ec.errors) > 0 {
		return nil, ec.errors
	}
	return injectorFiles, nil
}

type InjectorArgs struct {
	// Name is the name of the injector function.
	Name string
	// Tuple represents the arguments.
	Tuple *types.Tuple
	// Pos is the source position of the injector function.
	Pos token.Pos
}

func (g *gen) p(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// A wireErr is an error with an optional position.
type wireErr struct {
	error    error
	position token.Position
}

type outputSignature struct {
	out     types.Type
	cleanup bool
	err     bool
}

func injectorFuncSignature(sig *types.Signature) (*types.Tuple, outputSignature, error) {
	out, err := funcOutput(sig)
	if err != nil {
		return nil, outputSignature{}, err
	}
	return sig.Params(), out, nil
}

var (
	errorType   = types.Universe.Lookup("error").Type()
	cleanupType = types.NewSignature(nil, nil, nil, false)
)

// funcOutput validates an injector or provider function's return signature.
func funcOutput(sig *types.Signature) (outputSignature, error) {
	results := sig.Results()
	switch results.Len() {
	case 0:
		return outputSignature{}, errors.New("no return values")
	case 1:
		return outputSignature{out: results.At(0).Type()}, nil
	case 2:
		out := results.At(0).Type()
		switch t := results.At(1).Type(); {
		case types.Identical(t, errorType):
			return outputSignature{out: out, err: true}, nil
		case types.Identical(t, cleanupType):
			return outputSignature{out: out, cleanup: true}, nil
		default:
			return outputSignature{}, fmt.Errorf("second return type is %s; must be error or func()", types.TypeString(t, nil))
		}
	case 3:
		if t := results.At(1).Type(); !types.Identical(t, cleanupType) {
			return outputSignature{}, fmt.Errorf("second return type is %s; must be func()", types.TypeString(t, nil))
		}
		if t := results.At(2).Type(); !types.Identical(t, errorType) {
			return outputSignature{}, fmt.Errorf("third return type is %s; must be error", types.TypeString(t, nil))
		}
		return outputSignature{
			out:     results.At(0).Type(),
			cleanup: true,
			err:     true,
		}, nil
	default:
		return outputSignature{}, errors.New("too many return values")
	}
}
