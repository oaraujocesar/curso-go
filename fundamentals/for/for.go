package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// standard for
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d, ", i)

	}

	// infinity
	for {
		fmt.Println("4evah")
		time.Sleep(time.Second)
	}
}
