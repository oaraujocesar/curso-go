package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car
	isTurboEnabled bool
}

func main() {
	f := ferrari{}

	f.name = "F40"
	f.currentSpeed = 0
	f.isTurboEnabled = true

	fmt.Printf("is the Ferrari %s with turbo enabled? %v\n", f.name, f.isTurboEnabled)
}
