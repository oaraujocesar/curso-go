package main

import (
	"fmt"

	"github.com/oaraujocesar/area"
)

func main() {
	fmt.Printf("area: %.2f\n", area.Circle(6.0))
	fmt.Printf("area: %.2f\n", area.Rectangle(6.0, 3.0))
	fmt.Printf("Pi: %.2f\n", area.Pi)

	// this function is only accessible inside its own package.
	// fmt.Printf("area: %.2f\n", area._TriangleEq(5.0, 2.0))
}
