package main

import "fmt"

type Coffee interface {
	Cost() float64
}

type SimpleCoffee struct{}

func (coffee *SimpleCoffee) Cost() float64 {
	return 0.5
}

type TeaCoffee struct {
	coffee Coffee
}

func (coffee *TeaCoffee) Cost() float64 {
	return coffee.coffee.Cost() + 0.5
}

type MilkCoffee struct {
	coffee Coffee
}

func (coffee *MilkCoffee) Cost() float64 {
	return coffee.coffee.Cost() + 1.0
}

func Simple() {
	simpleCoffee := SimpleCoffee{}

	teaCoffee := TeaCoffee{
		coffee: &simpleCoffee,
	}

	milkCoffee := MilkCoffee{
		coffee: &simpleCoffee,
	}

	fmt.Println("Result of tea cost %v:", teaCoffee.Cost())
	fmt.Println("Result of milk cost %v:", milkCoffee.Cost())
}
