package main

import (
	"testing"

	"party/greet"
	mock_greet "party/greet/mock"

	"github.com/golang/mock/gomock"
)

func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		visitorLister *mock_greet.MockVisitorLister
		greeter       *mock_greet.MockGreeter
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
			args: args{
				justNice: false,
			},
			prepare: func(fileds *fields, args *args) {
				gomock.InOrder(
					fileds.visitorLister.EXPECT().ListVisitors(gomock.Any()).Return([]greet.Visitor{{"Peter", "TheSmart"}}, nil).Times(2),
					fileds.greeter.EXPECT().Hello(gomock.Any()).AnyTimes())
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tt.fields.visitorLister = mock_greet.NewMockVisitorLister(ctrl)
		tt.fields.greeter = mock_greet.NewMockGreeter(ctrl)

		if tt.prepare != nil {
			tt.prepare(&tt.fields, &tt.args)
		}

		t.Run(tt.name, func(t *testing.T) {
			s := &PartyService{
				visitorLister: tt.fields.visitorLister,
				greeter:       tt.fields.greeter,
			}
			if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
				t.Errorf("PartyService.GreetVisitors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
