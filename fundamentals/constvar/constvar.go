package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var ray = 3.2

	area := PI * m.Pow(ray, 2)

	fmt.Printf("The area calculated is: %.2fm2", area)
}
