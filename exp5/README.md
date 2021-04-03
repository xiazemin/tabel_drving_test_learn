https://www.programmersought.com/article/79944747770/

https://pkg.go.dev/golang.org/x/tools/go/packages


https://github.com/google/wire

% go run main.go load/tpkg
-: package load/tpkg is not in GOROOT (/usr/local/go/src/load/tpkg)
exit status 1


% go run main.go golang.org/x/tools/go/packages
golang.org/x/tools/go/packages [/Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/doc.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/external.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/golist.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/golist_overlay.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/loadmode_string.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/packages.go /Users/xiazemin/go/pkg/mod/golang.org/x/tools@v0.1.0/go/packages/visit.go]

% go run main.go github.com/labstack/echo
-: no required module provides package github.com/labstack/echo; to add it:
        go get github.com/labstack/echo
exit status 1


 % go get github.com/labstack/echo/v4
go get: added github.com/labstack/echo/v4 v4.2.1

% go run main.go github.com/labstack/echo/v4
github.com/labstack/echo/v4 [/Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/bind.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/binder.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/context.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/echo.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/group.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/ip.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/log.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/response.go /Users/xiazemin/go/pkg/mod/github.com/labstack/echo/v4@v4.2.1/router.go]


% go run main.go load/testpkg               
load/testpkg [/Users/xiazemin/source/tabel_drving_test_learn/exp5/testpkg/test.go]



https://scene-si.org/2018/06/19/listing-interfaces-with-go-ast-for-gomock-moq/

https://stackoverflow.com/questions/58906430/use-package-go-parser-to-read-ast-of-go-source-code-to-get-an-interfaces-method
https://stackoverflow.com/questions/33836358/parsing-go-src-trying-to-convert-ast-gendecl-to-types-interface

https://developers.mattermost.com/blog/instrumenting-go-code-via-ast-2/

% go run main.go load/testpkg
load/testpkg Mytpkg [/Users/xiazemin/source/tabel_drving_test_learn/exp5/testpkg/test.go]
9999-------- load/testpkg
[0x14000198080] pkg id name  path: load/testpkg Mytpkg load/testpkg
&{<nil> 17 type 0 [0x14000182480] 0}
interface: inter method: Hello
&{<nil> <nil> tpkg 0x140001b82a0 0x140001825d0}
<nil> tpkg &{67 0x14000182570 <nil>}
&{<nil> <nil> add 0x140001b82b8 0x14000182720}
<nil> add &{85 0x14000182660 0x14000182690}

golang.org/x/tools/internal/imports

https://nikodoko.com/posts/goimports_explained/