package greet

import (
	"fmt"
)

//Visitor ...
type Visitor struct {
	Name    string
	Surname string
}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}

//VisitorGroup ...
type VisitorGroup string

const (
	//NiceVisitor ...
	NiceVisitor VisitorGroup = "nice"
	//NotNiceVisitor ...
	NotNiceVisitor VisitorGroup = "not-nice"
)

//Greeter ...
type Greeter interface {
	Hello(name string) string
}

//VisitorLister ...
type VisitorLister interface {
	ListVisitors(who VisitorGroup) ([]Visitor, error)
}
