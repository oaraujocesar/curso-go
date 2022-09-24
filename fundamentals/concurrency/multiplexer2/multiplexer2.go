package main

import (
	"fmt"
	"time"
)

// generator
func speak(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)

			c <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}()

	return c
}

func group(firstEntry, secondEntry <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-firstEntry:
				c <- s
			case s := <-secondEntry:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := group(speak("John"), speak("Mary"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
