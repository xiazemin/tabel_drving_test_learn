package main

import "testing"

func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		visitorLister VisitorLister
		greeter       Greeter
	}
	type args struct {
		justNice bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
