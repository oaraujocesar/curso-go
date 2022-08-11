package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	firstName string
	lastName  string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p product) toString() string {
	return fmt.Sprintf("%s - US$:%2.f", p.name, p.price)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	var thing printable = person{firstName: "César", lastName: "Araújo"}

	fmt.Println(thing.toString())

	var secondThing printable = product{name: "Ferrari", price: 1699000.0}
	fmt.Println(secondThing.toString())

	thing = product{name: "Toy", price: 3.99}

	print(thing)
	print(secondThing)
}
