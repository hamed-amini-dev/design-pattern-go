package main

import "fmt"

type Iplace interface {
	Access(IVisitor)
}

type IVisitor interface {
	VisitShop(i *iShop)
	VisitFactory(i *iFactory)
}

type iShop struct {
}

func (is *iShop) Access(v IVisitor) {
	v.VisitShop(is)
}

type iFactory struct {
}

func (is *iFactory) Access(v IVisitor) {
	v.VisitFactory(is)
}

type iTaxVisit struct {
}

func (it *iTaxVisit) VisitShop(i *iShop) {
	fmt.Println("Shop Tax")
}

func (it *iTaxVisit) VisitFactory(i *iFactory) {
	fmt.Println("Factory Tax")
}

func main() {
	shop := &iShop{}
	factory := &iFactory{}

	auditor := &iTaxVisit{}

	shop.Access(auditor)
	factory.Access(auditor)
}
