package main

import "fmt"

// This example demonstrates "inheritance" (composition) in Go,
// as well as the Liskov Substitution Principle.
// Types that implement the AnimalBehavior interface are interchangeable,
// without affecting the correctness of the program, the MakeSound func

type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Printf("Animal sound: %s\n", a.Name)
}

type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Printf("Chirp chirp: %s\n", b.Name)
}

type AnimalBehavior interface {
	MakeSound()
}

func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}

func main() {
	a := Animal{Name: "Hippo"}
	b := Bird{Animal{Name: "Starling"}}
	MakeSound(a)
	MakeSound(b)
}
