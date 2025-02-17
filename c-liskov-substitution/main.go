package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}

type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

type AnimalBehavior interface {
	MakeSound()
}

// This demonstrates inheritance in Go,
// as well as the Liskov Substitution Principle.
// Objects of a subtype Bird can be used wherever objects
// of the base type Animal are expected,
// without affecting the correctness of the program

func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

func main() {
	a := Animal{}
	b := Bird{}
	MakeSound(a)
	MakeSound(b)
}
