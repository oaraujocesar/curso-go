package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4                  // buffer fullfilled
	fmt.Println("executed!") // not executed
	ch <- 5
	ch <- 6
	ch <- 7
}

func main() {
	ch := make(chan int, 3)

	go routine(ch)

	time.Sleep(time.Second * 3)
	fmt.Println(<-ch)
}
