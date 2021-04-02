package main

type greeter struct {
}

func (g greeter) Hello(name string) string {
	return name
}

type visitorLister struct {
}

func (v visitorLister) ListVisitors(who VisitorGroup) ([]Visitor, error) {
	return []Visitor{
		{
			Name: "test",
		},
	}, nil
}
