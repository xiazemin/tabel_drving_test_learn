package main

import (
	"fmt"
	"party/greet"
)

//PartyService ...
type PartyService struct {
	visitorLister greet.VisitorLister
	greeter       greet.Greeter
}

//NewPartyService ...
func NewPartyService(namesService greet.VisitorLister, greeter greet.Greeter) *PartyService {
	return &PartyService{
		visitorLister: namesService,
		greeter:       greeter,
	}
}

//GreetVisitors ...
func (s *PartyService) GreetVisitors(justNice bool) error {
	visitors, err := s.visitorLister.ListVisitors(greet.NiceVisitor)
	if err != nil {
		return fmt.Errorf("could get nice people names: %w", err)
	}
	if !justNice {
		notNice, err := s.visitorLister.ListVisitors(greet.NotNiceVisitor)
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

//GreetVisitors ...
func (s *PartyService) GreetVisitors1(justNice bool) error {
	visitors, err := s.visitorLister.ListVisitors(greet.NiceVisitor)
	if err != nil {
		return fmt.Errorf("could get nice people names: %w", err)
	}
	if !justNice {
		notNice, err := s.visitorLister.ListVisitors(greet.NotNiceVisitor)
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
