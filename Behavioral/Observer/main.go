package main

import "fmt"

type Observable interface {
	Update(ID int) error
}

type ConcreteObservable struct {
	Name string
	Id   int
}

func (o *ConcreteObservable) Update(ID int) error {
	o.Id = ID
	return nil
}

type Subjects struct { // like app
	observable []Observable // list observer
	state      int
}

func (s *Subjects) Attach(ob Observable) {
	s.observable = append(s.observable, ob)
}

func (s *Subjects) Notify(state int) error {
	s.state = state
	for _, ob := range s.observable {
		if err := ob.Update(state); err != nil {
			return err
		}
	}
	return nil
}

func (s *Subjects) GetIDs() {
	for _, ob := range s.observable {
		fmt.Printf("Onservable %s is %d\n", ob.(*ConcreteObservable).Name, ob.(*ConcreteObservable).Id)
	}
}

func main() {
	co1 := ConcreteObservable{
		Name: "OB0",
		Id:   0,
	}
	co2 := ConcreteObservable{
		Name: "OB1",
		Id:   0,
	}

	sub := Subjects{state: 0}

	sub.Attach(&co1)
	sub.Attach(&co2)
	//
	err := sub.Notify(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	sub.GetIDs()
}
