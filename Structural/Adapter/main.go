package main

import "fmt"

type IOldPrinter interface {
	OldPrint(string)
}

type OldPrinter struct {
}

func (p *OldPrinter) OldPrint(s string) {
	fmt.Println(s)
}

type NewPrinter interface {
	Print(string)
}

type PrinterAdapter struct {
	OldPrinter IOldPrinter
}

func (p *PrinterAdapter) Print(s string) {
	p.OldPrinter.OldPrint(s)
}

func main() {
	oldPrinter := &OldPrinter{}
	newPrinter := &PrinterAdapter{oldPrinter}
	newPrinter.Print("Hello World")
}
