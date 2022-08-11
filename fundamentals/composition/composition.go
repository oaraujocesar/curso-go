package main

import "fmt"

type sport interface {
	enableTurbo()
}

type deluxe interface {
	autoParking()
}

type luxurySport interface {
	sport
	deluxe
}

type bmw7 struct{}

func (b bmw7) enableTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) autoParking() {
	fmt.Println("Parking...")
}

func main() {
	var b luxurySport = bmw7{}
	b.autoParking()
	b.enableTurbo()
}
