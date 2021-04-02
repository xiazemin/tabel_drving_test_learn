#问题1: lint 不通过
% golangci-lint run .
party_test.go:24:20: Using the variable on range scope `tt` in function literal (scopelint)
                                visitorLister: tt.fields.visitorLister,
                                               ^
party_test.go:25:20: Using the variable on range scope `tt` in function literal (scopelint)
                                greeter:       tt.fields.greeter,
                                               ^
party_test.go:27:30: Using the variable on range scope `tt` in function literal (scopelint)
                        if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
                                                  ^
party_test.go:28:74: Using the variable on range scope `tt` in function literal (scopelint)
                                t.Errorf("PartyService.GreetVisitors() error = %v, wantErr %v", err, tt.wantErr)

#问题2: 传的context不对，会导致逻辑异常，比如从context里面获取数据，一个case应该有统一的context，

#问题3: 全局的gomock导致case间相互影响，特别是.AnyTimes,应该有隔离的mockerController

