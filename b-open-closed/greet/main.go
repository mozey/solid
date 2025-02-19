package main

import "fmt"

// Demonstrates how types are open for extension by embedding

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello GolangUK", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to Go UK", b.year)
}

func main() {
	var a A
	a.year = 2016
	b := B{A{year: 2017}}
	a.Greet()
	b.Greet()
}
