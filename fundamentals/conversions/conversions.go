package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	result := 6.9
	finalResult := int(result)

	fmt.Println(finalResult)

	// this evaluates to a character from unicode table
	fmt.Println("test " + string(123))

	// int to string
	fmt.Println("Test", strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 3)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("It works!")
	}
}
