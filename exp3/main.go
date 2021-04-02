package main

func main() {

	p := NewPartyService(visitorLister{}, greeter{})
	_ = p.GreetVisitors(true)

}
