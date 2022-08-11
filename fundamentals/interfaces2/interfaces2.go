package main

import "fmt"

type sport interface {
	enableTurbo()
}

type ferrari struct {
	model          string
	isTurboEnabled bool
	currentSpeed   int
}

func (f *ferrari) enableTurbo() {
	f.isTurboEnabled = true
}

func main() {
	firstCar := ferrari{model: "F40", isTurboEnabled: false, currentSpeed: 0}
	firstCar.enableTurbo()

	var secondCar sport = &ferrari{model: "F50", isTurboEnabled: false, currentSpeed: 0}
	secondCar.enableTurbo()

	fmt.Println(firstCar, secondCar)
}
