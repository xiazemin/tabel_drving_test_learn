package main

import (
	"party/greet"
	"testing"

	mock_greet "party/greet/mock"

	"github.com/golang/mock/gomock"
)

func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		visitorLister *mock_greet.MockVisitorLister

		greeter *mock_greet.MockGreeter
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
			prepare: func(fields *fields, args *args) {
				gomock.InOrder(

					fields.visitorLister.EXPECT().ListVisitors(greet.VisitorGroup{}).Return([]greet.Visitor{},
						error,
					).Times(2),

					fields.visitorLister.EXPECT().ListVisitors2(&greet.VisitorGroup{}).Return([]*greet.Visitor{},
						error,
					).Times(2),

					fields.greeter.EXPECT().Hello(string).Return(string).Times(2),

					fields.greeter.EXPECT().Hello1(string,
						string,
					).Return(string,
						string,
					).Times(2),
				)
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
