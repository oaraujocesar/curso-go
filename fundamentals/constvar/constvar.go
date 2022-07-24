package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var ray = 3.2

	area := PI * m.Pow(ray, 2)

	fmt.Printf("The calculated area is: %.2fm2 \n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Oops!"

	fmt.Println(g, h, i)
}
