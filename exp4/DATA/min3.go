Generated TestPartyService_GreetVisitors
package main

import (
	"party/greet"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		visitorLister greet.VisitorLister
		greeter       greet.Greeter
	}
	type args struct {
		justNice bool
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(fields *fields, args *args)
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			prepare: func(fileds *fields, args *args) {
				gomock.InOrder(
					fileds.visitorLister.EXPECT().ListVisitors(gomock.Any()).Return([]greet.Visitor{&greet.Visitor{"Peter", "TheSmart"}}, nil).Times(2),
					fileds.greeter.EXPECT().Hello(gomock.Any()).AnyTimes())
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tt.fields.visitorLister = mock_greet.NewMockVisitorLister(ctrl)
		tt.fields.greeter = mock_greet.NewMockGreeter(ctrl)
		if tt.prepare != nil {
			tt.prepare(&tt.fields, &tt.args)
		}
		s := &PartyService{
			visitorLister: tt.fields.visitorLister,
			greeter:       tt.fields.greeter,
		}
		if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
			t.Errorf("%q. PartyService.GreetVisitors() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
