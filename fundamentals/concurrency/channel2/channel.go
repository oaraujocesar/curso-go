package main

import (
	"fmt"
	"time"
)

func multiply(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // writing to the channel

	time.Sleep(time.Second * 2)
	c <- 3 * base // writing to the channel

	time.Sleep(time.Second * 3)
	c <- 4 * base // writing to the channel
}

func main() {
	ch := make(chan int)

	go multiply(2, ch)

	a, b := <-ch, <-ch // reading data from the channel

	fmt.Println(a, b)
	fmt.Println(<-ch)
}
