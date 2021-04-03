go get github.com/golang/mock/mockgen

mockgen -source=greet/greeter.go -destination=greet/mock/greeter.go

go test ./

#gotests -only "函数名称" 文件名称.go

gotests -only GreetVisitors /Users/xiazemin/software/tabel_drving_test_learn/exp4/party.go 


gotests -only GreetVisitors party.go > party_test.go


% gotests -w -only ListVisitors greeter.go                
Generated Test_visitorLister_ListVisitors

% gotests -w -only GreetVisitors party.go                 
Generated TestPartyService_GreetVisitors