package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) getFullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *person) setName(name string) {
	splitedName := strings.Split(name, " ")

	p.firstName = splitedName[0]
	p.lastName = splitedName[1]
}

func main() {

	p1 := person{
		firstName: "César",
		lastName:  "Araújo",
	}

	fmt.Println(p1.getFullName())

	p1.setName("Francielly Kelly")

	fmt.Println(p1.getFullName())
}
