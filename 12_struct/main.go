package main

import (
	"fmt"
)

func main() {

	fmt.Println("halo zl")

	var animal Animal
	animal = &Dog{}
	fmt.Println("animal name: ", animal.Name())
	fmt.Println("animal bark: ", animal.Bark())

	gd := &GuideDog{}
	gd.Help()
}

type PartyAnimal interface {
	Animal
	Party()
}

type Animal interface {
	Name() string
	Bark() string
}

type Dog struct {
	// sth
}

func (d *Dog) Name() string {
	// do sth
	return "Dog"
}

func (d *Dog) Bark() string {
	// do sth
	return "woof!"
}

type GuideDog struct {
	// do sth
	*Dog
}

func (gd *GuideDog) Help() {
	// do sth
	fmt.Printf("hey human, grab %s's leash\n", gd.Name())
}
