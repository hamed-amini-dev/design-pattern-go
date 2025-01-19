package main

import "fmt"

// implementation
type IDevice interface {
	on()
	off()
}

// abstract
type IRemoteDevice struct {
	dv IDevice
}

// concreate implementaion
type iLight struct {
}

func (i *iLight) on() {
	fmt.Println("Light is on")
}
func (i *iLight) off() {
	fmt.Println("Light is off")
}

func main() {
	il := &iLight{}
	IR := IRemoteDevice{
		dv: il,
	}

	IR.dv.on()
	IR.dv.off()

}
