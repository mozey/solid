package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const RoleAdmin = "admin"

type User struct {
	// role is private, and can't be edited
	role string
}

// NewUser constructor
func NewUser(role string) User {
	return User{role: role}
}

func (u User) Role() string {
	return u.role
}

type EmployeeInfo struct {
	// Name is public, don't create unnecessary getter methods
	Name string
	// salary is private
	salary float64
}

// NewEmployeeInfo constructor
func NewEmployeeInfo(name string, salary float64) EmployeeInfo {
	return EmployeeInfo{Name: name, salary: salary}
}

func (e EmployeeInfo) Salary(u User) (salary float64, err error) {
	if u.Role() != "admin" {
		// return errors.Errorf
	}
	return e.salary, nil
}

type Address struct {
	line1 string
	line2 string
	line3 string
}

// NewAddress constructor
func NewAddress(line1, line2, line3 string) Address {
	return Address{line1: line1, line2: line2, line3: line3}
}

func (e Address) Address() string {
	return fmt.Sprintf("%s, %s, %s", e.line1, e.line2, e.line3)
}

type Employee struct {
	Name    string
	Salary  float64
	Address string
}

type NewEmployeeParams struct {
	User    User
	Info    EmployeeInfo
	Address Address
}

// NewEmployee constructor func
func NewEmployee(params NewEmployeeParams) (e Employee, err error) {
	salary, err := params.Info.Salary(params.User)
	if err != nil {
		return e, err
	}
	return Employee{
		Name:    params.Info.Name,
		Salary:  salary,
		Address: params.Address.Address(),
	}, nil
}

func (e Employee) JSON() (b []byte, err error) {
	type wrapper struct{ Employee Employee }
	b, err = json.MarshalIndent(wrapper{e}, "", "  ")
	if err != nil {
		return b, err
	}
	return b, nil
}

func main() {
	// Session user
	u := NewUser(RoleAdmin)

	// Info
	i := NewEmployeeInfo("Joe", 123.0)

	// Address
	a := NewAddress(i.Name, "4 Sunny Str", "South Africa")

	// Employee
	e, err := NewEmployee(
		// Params use composition
		NewEmployeeParams{
			User:    u,
			Info:    i,
			Address: a,
		})

	// Print JSON
	b, err := e.JSON()
	if err != nil {
		fmt.Sprintln(err.Error())
		os.Exit(1)
	}
	fmt.Printf("\n%s\n\n", b)
}
