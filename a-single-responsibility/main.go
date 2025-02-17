package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const RoleAdmin = "admin"

type UserClaims struct {
	// role is private, and can't be edited
	role string
}

// NewUserClaims constructor
func NewUserClaims(role string) UserClaims {
	return UserClaims{role: role}
}

func (u UserClaims) Role() string {
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

func (e EmployeeInfo) Salary(u UserClaims) (salary float64, err error) {
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
	Claims  UserClaims
	Info    EmployeeInfo
	Address Address
}

// NewEmployee constructor func
func NewEmployee(params NewEmployeeParams) (e Employee, err error) {
	salary, err := params.Info.Salary(params.Claims)
	if err != nil {
		salary = 0 // Set default value
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
	// UserClaims, i.e. parsed from a JWT for a session
	u := NewUserClaims(RoleAdmin)

	// Info
	i := NewEmployeeInfo("Joe", 123.0)

	// Address
	a := NewAddress(i.Name, "4 Sunny Str", "South Africa")

	// Employee
	e, err := NewEmployee(
		// Params use composition
		NewEmployeeParams{
			Claims:  u,
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
