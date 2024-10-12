package main

import "fmt"

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Printf("Storing  %s\n", value)
	return nil
}

type ExecuteFN func(string) error

func Exec(db DB) ExecuteFN {
	return func(value string) error {
		fmt.Printf("Executing %s\n", value)
		return db.Store(value)
	}
}

func Execute(fn ExecuteFN) error {
	err := fn("Hello World")
	if err != nil {
		return err
	}
	return nil
}

func DBSample() {
	s := &Store{}
	err := Execute(Exec(s))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
