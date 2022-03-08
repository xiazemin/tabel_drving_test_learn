package main

import (
	"party/greet"
)

type greeter struct {
}

func (g greeter) Hello(name string) string {
	return name
}
func (g greeter) Hello1(name, name2 string) (string, string) {
	return name, name2
}

type visitorLister struct {
}

func (v visitorLister) ListVisitors(who greet.VisitorGroup) ([]greet.Visitor, error) {
	return []greet.Visitor{
		{
			Name: "test",
		},
	}, nil
}

func (v visitorLister) ListVisitors2(who *greet.VisitorGroup) ([]*greet.Visitor, error) {
	return []*greet.Visitor{
		{
			Name: "test",
		},
	}, nil
}
