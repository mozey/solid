package main

import (
	"fmt"
	"log"
	"os"
)

// Testing the code below is difficult,
// customerService relies on the actual implementation to store a Customer.
// See "Decoupling" example here: https://100go.co/5-interface-pollution/

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

type MySQL struct{}

func (storer MySQL) StoreCustomer(c Customer) error {
	fmt.Printf("Created a new customer with ID %s\n", c.id)
	return nil
}

func (storer MySQL) CreateCustomerQuery(c Customer) string {
	return fmt.Sprintf("Query to create customer %s", c.id)
}

type QueryLogger struct {
	MySQL
}

func (storer QueryLogger) StoreCustomer(c Customer) error {
	log.Print(storer.CreateCustomerQuery(c))
	storer.MySQL.StoreCustomer(c)
	return nil
}

func main() {
	s := MySQL{}
	cs := CustomerService{storer: s}
	err := cs.CreateNewCustomer("123")
	if err != nil {
		os.Exit(1)
	}

	// The storer base type (MySQL) is interchangeable
	// with the subtype (QueryLogger)
	l := QueryLogger{}
	cs = CustomerService{storer: l}
	err = cs.CreateNewCustomer("456")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
