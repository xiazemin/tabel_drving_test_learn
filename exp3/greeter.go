package main

import (
	"party/greet"
)

type greeter struct {
}

func (g greeter) Hello(name string) string {
	return name
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
