import (
	"party/greet"
	"testing"
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
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &PartyService{
			visitorLister: tt.fields.visitorLister,
			greeter:       tt.fields.greeter,
		}
		if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
			t.Errorf("%q. PartyService.GreetVisitors() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}