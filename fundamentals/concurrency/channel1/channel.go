package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 2 // write

	fmt.Println(<-ch) // read
}
