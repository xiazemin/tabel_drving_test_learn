package main

import (
	"testing"

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
		name   string
		fields fields
		// prepare lets us initialize our mocks within the `tests` slice. Oftentimes it also proves useful for other initialization
		prepare func(fields *fields, args *args)
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				justNice: true,
			},
			prepare: func(fileds *fields, args *args) {
				// mockedVisitorLister called once with party.NiceVisitor argument would return []string{"Peter"}, nil
				//fileds.visitorLister.EXPECT().ListVisitors(greet.NiceVisitor).Return([]greet.Visitor{0: greet.Visitor{"Peter", "TheSmart"}}, nil)
				fileds.greeter.EXPECT().Hello(gomock.Any())
				// mockedVisitorLister called once with party.NotNiceVisitor argument would return nil and error
				//mockedVisitorLister.EXPECT().ListVisitors(greet.NotNiceVisitor).Return(nil, fmt.Errorf("dummyErr"))
				// mockedVisitorLister implements party.VisitorLister interface, so it can be assigned in PartyService
			},
			wantErr: false,
		},
		{
			name: "case2",
			args: args{
				justNice: false,
			},
			prepare: func(fileds *fields, args *args) {
				// mockedVisitorLister called once with party.NiceVisitor argument would return []string{"Peter"}, nil
				gomock.InOrder(
					//fileds.visitorLister.EXPECT().ListVisitors(gomock.Any()).Return([]greet.Visitor{0: greet.Visitor{"Peter", "TheSmart"}}, nil).Times(2),
					fileds.greeter.EXPECT().Hello(gomock.Any()).AnyTimes())
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		ctrl := gomock.NewController(t)
		// if not all expectations set on the controller are fulfilled at the end of function, the test will fail!
		defer ctrl.Finish()
		// init structure which implements party.NamesLister interface
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
