package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i

	fmt.Println(p)  // or &p: the address
	fmt.Println(*p) // the value

	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
