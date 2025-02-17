package main

import "fmt"

// Testing the code below is difficult,
// customerService relies on the actual implementation to store a Customer

// type CustomerService struct {
//     store mysql.Store // Depends on the concrete implementation
// }

// func (cs CustomerService) CreateNewCustomer(id string) error {
//     customer := Customer{id: id}
//     return cs.store.StoreCustomer(customer)
// }

// Decouple CustomerService from the actual implementation:

type Customer struct {
	id string
}

type customerStorer interface { // Creates a storage abstraction
	StoreCustomer(Customer) error
}

type CustomerService struct {
	// storer decouples CustomerService from the actual implementation
	storer customerStorer
}

func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.storer.StoreCustomer(customer)
}

type Logger struct{}

func (l Logger) StoreCustomer(c Customer) error {
	fmt.Printf("\nCustomer ID is %s\n\n", c.id)
	return nil
}

func main() {
	l := Logger{}
	c := Customer{id: "customer123"}
	cs := CustomerService{storer: l}
	cs.storer.StoreCustomer(c)
}
