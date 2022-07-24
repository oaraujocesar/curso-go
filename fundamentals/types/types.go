package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// int
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal integer is:", reflect.TypeOf(32000))

	// uint8, 16, 32, 64
	var b byte = 255
	fmt.Println("The byte is:", reflect.TypeOf(b))

	// int8, 16, 32, 64
	i1 := math.MaxInt64
	fmt.Println(i1, reflect.TypeOf(i1))

	var i2 rune = 'a' // represents a value from unicode table (int32)
	fmt.Println(i2, reflect.TypeOf(i2))

	// float-point numbers
	var x float32 = 49.99
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello, my name is César"
	fmt.Println(s1 + "!")
	fmt.Println("The string's length is:", len(s1))

	// multiline string
	s2 := `Hello,
my
name
is
César`
	fmt.Println(s2)
	fmt.Println("The string's length is:", len(s2))
}
