go get github.com/golang/mock/mockgen

mockgen -source=greet/greeter.go -destination=greet/mock/greeter.go

go test ./