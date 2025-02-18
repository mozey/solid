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

type DB struct{}

func (storer DB) StoreCustomer(c Customer) error {
	fmt.Printf("Created a new customer with ID %s\n", c.id)
	return nil
}

func (storer DB) CreateCustomerQuery(c Customer) string {
	return fmt.Sprintf("Query to create customer %s", c.id)
}

type QueryLogger struct {
	DB
}

func (storer QueryLogger) StoreCustomer(c Customer) error {
	log.Print(storer.CreateCustomerQuery(c))
	storer.DB.StoreCustomer(c)
	return nil
}

func main() {
	s := DB{}
	cs := CustomerService{storer: s}
	err := cs.CreateNewCustomer("123")
	if err != nil {
		os.Exit(1)
	}

	l := QueryLogger{}
	cs = CustomerService{storer: l}
	err = cs.CreateNewCustomer("456")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
