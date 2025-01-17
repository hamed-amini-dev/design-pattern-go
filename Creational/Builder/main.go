package main

import "fmt"

type House struct {
	Windows int
	Doors   int
}

type IHouseBuilder interface {
	SetWindows(int) IHouseBuilder
	SetDoors(int) IHouseBuilder
	Build() House
}

type iHouse struct {
	Windows int
	Doors   int
}

func NewiHouse() IHouseBuilder {
	return &iHouse{}
}

func (h *iHouse) SetWindows(w int) IHouseBuilder {
	h.Windows = w
	return h
}

func (h *iHouse) SetDoors(d int) IHouseBuilder {
	h.Doors = d
	return h
}

func (h *iHouse) Build() House {
	return House{
		Windows: h.Windows,
		Doors:   h.Doors,
	}
}

func main() {
	h := NewiHouse()

	h.SetWindows(2).SetDoors(5).Build()

	fmt.Println(h)

}
