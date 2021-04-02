package main

import (
	"fmt"
	"testing"

	"party/greet"
	mock_greet "party/greet/mock"

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
	ctrl := gomock.NewController(t)
	// if not all expectations set on the controller are fulfilled at the end of function, the test will fail!
	defer ctrl.Finish()
	// init structure which implements party.NamesLister interface
	mockedVisitorLister := mock_greet.NewMockVisitorLister(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ccase1",
			fields: fields{
				visitorLister: mockedVisitorLister,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// mockedVisitorLister called once with party.NiceVisitor argument would return []string{"Peter"}, nil
		mockedVisitorLister.EXPECT().ListVisitors(greet.NiceVisitor).Return([]greet.Visitor{{"Peter", "TheSmart"}}, nil)
		// mockedVisitorLister called once with party.NotNiceVisitor argument would return nil and error
		mockedVisitorLister.EXPECT().ListVisitors(greet.NotNiceVisitor).Return(nil, fmt.Errorf("dummyErr"))
		// mockedVisitorLister implements party.VisitorLister interface, so it can be assigned in PartyService

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
