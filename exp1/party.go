package main

import "fmt"

//VisitorGroup ...
type VisitorGroup string

const (
	//NiceVisitor ...
	NiceVisitor VisitorGroup = "nice"
	//NotNiceVisitor ...
	NotNiceVisitor VisitorGroup = "not-nice"
)

//Visitor ...
type Visitor struct {
	Name    string
	Surname string
}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}

//Greeter ...
type Greeter interface {
	Hello(name string) string
}

//VisitorLister ...
type VisitorLister interface {
	ListVisitors(who VisitorGroup) ([]Visitor, error)
}

//PartyService ...
type PartyService struct {
	visitorLister VisitorLister
	greeter       Greeter
}

//NewPartyService ...
func NewPartyService(namesService VisitorLister, greeter Greeter) *PartyService {
	return &PartyService{
		visitorLister: namesService,
		greeter:       greeter,
	}
}

//GreetVisitors ...
func (s *PartyService) GreetVisitors(justNice bool) error {
	visitors, err := s.visitorLister.ListVisitors(NiceVisitor)
	if err != nil {
		return fmt.Errorf("could get nice people names: %w", err)
	}
	if !justNice {
		notNice, err := s.visitorLister.ListVisitors(NotNiceVisitor)
		if err != nil {
			return fmt.Errorf("could not get not-nice people's names' ")
		}
		visitors = append(visitors, notNice...)
	}
	for _, visitor := range visitors {
		fmt.Println(s.greeter.Hello(visitor.String()))
	}
	return nil
}
