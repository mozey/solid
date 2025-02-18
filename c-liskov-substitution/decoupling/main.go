package main

import (
	"fmt"
	"log"
	"os"
)

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

type CustomerStorer interface { // Creates a storage abstraction
	StoreCustomer(Customer) error
}

type CustomerService struct {
	// storer decouples CustomerService from the actual implementation
	storer CustomerStorer
}

func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.storer.StoreCustomer(customer)
}

type Logger struct{}

func (l Logger) StoreCustomer(c Customer) error {
	log.Printf("Created a new customer with ID %s", c.id)
	return nil
}

type Printer struct{}

func (l Printer) StoreCustomer(c Customer) error {
	fmt.Printf("Created a new customer with ID %s\n", c.id)
	return nil
}

func main() {
	l := Logger{}
	csLogger := CustomerService{storer: l}
	err := csLogger.CreateNewCustomer("123")
	if err != nil {
		os.Exit(1)
	}

	p := Printer{}
	csPrinter := CustomerService{storer: p}
	err = csPrinter.CreateNewCustomer("456")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
