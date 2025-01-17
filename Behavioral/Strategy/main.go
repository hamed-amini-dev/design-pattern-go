package main

import "fmt"

type IPaymentStrategy interface {
	Pay(amount float64)
}

type Stripe struct {
}

func (s *Stripe) Pay(amount float64) {
	fmt.Println("amount :", amount)
}

type PayPal struct {
}

func (s *PayPal) Pay(amount float64) {
	fmt.Println("amount :", amount)
}

type PaymentStrategy struct {
	strategy IPaymentStrategy
}

func (p *PaymentStrategy) SetStrategy(s IPaymentStrategy) {
	p.strategy = s
}

func (p *PaymentStrategy) PayBill(amount float64) {
	p.strategy.Pay(amount)
}

func main() {
	stripe := &Stripe{}

	strategy := PaymentStrategy{}

	strategy.SetStrategy(stripe)

	strategy.PayBill(64)

}
