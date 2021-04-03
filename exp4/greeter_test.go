package main

import (
	"party/greet"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_visitorLister_ListVisitors(t *testing.T) {
	type args struct {
		who greet.VisitorGroup
	}
	tests := []struct {
		name    string
		v       visitorLister
		args    args
		want    []greet.Visitor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		tt := tt
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		v := visitorLister{}
		got, err := v.ListVisitors(tt.args.who)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. visitorLister.ListVisitors() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. visitorLister.ListVisitors() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
