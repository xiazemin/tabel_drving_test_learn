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
	Hello1(name1, name2 string) (string, string)
}

//VisitorLister ...
type VisitorLister interface {
	ListVisitors(who VisitorGroup) ([]Visitor, error)
	ListVisitors2(who *VisitorGroup) ([]*Visitor, error)
}
