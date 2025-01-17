package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Subscribers struct {
	observers []Observer
}

func (s *Subscribers) Add(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subscribers) Notify(message string) {
	for _, s := range s.observers {
		s.Update(message)
	}
}

type Subscrier struct {
	name string
}

func (s *Subscrier) Update(message string) {
	fmt.Println(s.name, message)
}

func main() {
	s := Subscribers{
		observers: make([]Observer, 0),
	}

	s.Add(&Subscrier{"A"})
	s.Add(&Subscrier{"B"})

	s.Notify("Hello")
}
