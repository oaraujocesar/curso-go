package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)

	c <- 1 // blocking operation

	fmt.Println("after being read")
}

func main() {
	c := make(chan int) // bufferless channel

	go routine(c)

	fmt.Println(<-c) // blocking operation
	fmt.Println("readed")
	fmt.Println(<-c) // deadlock! goroutines are dead.
	fmt.Println("End")
}
