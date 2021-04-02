package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "/tmp/tmp.go", `package service

    type ServiceInterface interface {
                    Create(NewServiceRequest) (JsonResponse, error)
                    Delete(DelServiceRequest) (JsonResponse, error)
    }`, 0)
	for _, x := range f.Decls {
		if x, ok := x.(*ast.GenDecl); ok {
			if x.Tok != token.TYPE {
				continue
			}
			for _, x := range x.Specs {
				if x, ok := x.(*ast.TypeSpec); ok {
					iname := x.Name
					if x, ok := x.Type.(*ast.InterfaceType); ok {
						for _, x := range x.Methods.List {
							if len(x.Names) == 0 {
								continue
							}
							mname := x.Names[0].Name
							fmt.Println("interface:", iname, "method:", mname)

						}
					}
				}
			}
		}
	}
}
