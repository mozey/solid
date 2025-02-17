package main

import "fmt"

type PaymentMethod interface {
	Pay()
}

// Payment struct is open for extension and closed for modification.
// The PaymentMethod interface is implemented when adding new payment methods

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard\n", cc.amount)
}

func main() {
	p := Payment{}
	cc := CreditCard{12.23}
	p.Process(cc)
}

// Adding something like PayPal would look like this...
