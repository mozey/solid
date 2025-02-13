package main

import "fmt"

// Better to split into EmployeeInfo and EmployeeAddress
// type Employee struct {
//     Name    string
//     Salary  float64
//     Address string
// }

type EmployeeInfo struct {
	Name   string
	Salary float64
}

type EmployeeAddress struct {
	Address string
}

func (e EmployeeInfo) GetSalary() float64 {
	return e.Salary
}

func (e EmployeeAddress) GetAddress() string {
	return e.Address
}

func main() {
	i := EmployeeInfo{Salary: 123}
	fmt.Println(i.GetSalary())

	a := EmployeeAddress{Address: "RSA"}
	fmt.Println(a.GetAddress())
}
