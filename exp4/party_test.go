package main

import (
	"party/greet"
	mock_greet "party/greet/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
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

func TestPartyService_GreetVisitors1(t *testing.T) {
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
			name:   "case1",
			fields: fields{},
			args: args{
				justNice: false, //bool
			},
			wantErr: false,
			prepare: func(fields *fields, args *args) {
				//httpmock.RegisterResponder("GET", "https://mytest.com/httpmock", httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Test"}]`))
				gomock.InOrder(
				//fields.visitorLister.EXPECT().ListVisitors(greet.NewVisitorGroup{}).Return([]greet.Visitor{}{},nil).AnyTimes(),  //params:greet.VisitorGroup{} ;  return:[]greet.Visitor{} error
				)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		httpmock.Activate()
		tt.fields.visitorLister = mock_greet.NewMockVisitorLister(ctrl)
		tt.fields.greeter = mock_greet.NewMockGreeter(ctrl)
		if tt.prepare != nil {
			tt.prepare(&tt.fields, &tt.args)
		}

		s := &PartyService{
			visitorLister: tt.fields.visitorLister,
			greeter:       tt.fields.greeter,
		}
		if err := s.GreetVisitors1(tt.args.justNice); (err != nil) != tt.wantErr {
			t.Errorf("%q. PartyService.GreetVisitors1() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
		httpmock.DeactivateAndReset()
	}
}
