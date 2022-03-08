package main

import (
	mock_greet "party/greet/mock"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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
		prepare func(fields *fields, args *args, mock redismock.ClientMock)
		args    args
		wantErr assert.ErrorAssertionFunc //bool
	}{
		// TODO: Add test cases.
		{
			name:   "case1",
			fields: fields{},
			args: args{
				justNice: false, //bool
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, assert.AnError, i)
			}, //如果没有error用assert.NoError,
			//false,
			prepare: func(fields *fields, args *args, mock redismock.ClientMock) {
				//httpmock.RegisterResponder("GET", "https://mytest.com/httpmock", httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Test"}]`))
				//mock.ExpectGet("prefix:10").RedisNil()
				// mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))
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
		//db, mock := redismock.NewClientMock()
		//把 db赋值给redis client github.com/go-redis/redis/v8
		_, mock := redismock.NewClientMock()
		tt.fields.visitorLister = mock_greet.NewMockVisitorLister(ctrl)
		tt.fields.greeter = mock_greet.NewMockGreeter(ctrl)
		if tt.prepare != nil {
			tt.prepare(&tt.fields, &tt.args, mock)
		}

		s := &PartyService{
			visitorLister: tt.fields.visitorLister,
			greeter:       tt.fields.greeter,
		}
		if err := s.GreetVisitors(tt.args.justNice); !tt.wantErr(t, err) { //(err != nil) != tt.wantErr
			t.Errorf("%q. PartyService.GreetVisitors() error = %v, wantErr %v", tt.name, err, tt.wantErr(t, err))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
		httpmock.DeactivateAndReset()
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
		prepare func(fields *fields, args *args, mock redismock.ClientMock)
		args    args
		wantErr assert.ErrorAssertionFunc //bool
	}{
		// TODO: Add test cases.
		{
			name:   "case1",
			fields: fields{},
			args: args{
				justNice: false, //bool
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, assert.AnError, i)
			}, //如果没有error用assert.NoError,
			//false,
			prepare: func(fields *fields, args *args, mock redismock.ClientMock) {
				//httpmock.RegisterResponder("GET", "https://mytest.com/httpmock", httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Test"}]`))
				//mock.ExpectGet("prefix:10").RedisNil()
				// mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))
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
		//db, mock := redismock.NewClientMock()
		//把 db赋值给redis client github.com/go-redis/redis/v8
		_, mock := redismock.NewClientMock()
		tt.fields.visitorLister = mock_greet.NewMockVisitorLister(ctrl)
		tt.fields.greeter = mock_greet.NewMockGreeter(ctrl)
		if tt.prepare != nil {
			tt.prepare(&tt.fields, &tt.args, mock)
		}

		s := &PartyService{
			visitorLister: tt.fields.visitorLister,
			greeter:       tt.fields.greeter,
		}
		if err := s.GreetVisitors1(tt.args.justNice); !tt.wantErr(t, err) { //(err != nil) != tt.wantErr
			t.Errorf("%q. PartyService.GreetVisitors1() error = %v, wantErr %v", tt.name, err, tt.wantErr(t, err))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
		httpmock.DeactivateAndReset()
	}
}
