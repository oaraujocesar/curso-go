package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int, c chan int) {
	start := 2

	for i := 0; i < n; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime

				start = prime + 1

				time.Sleep(time.Millisecond * 100)

				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 30)

	go primes(cap(c), c)

	for prime := range c {
		fmt.Printf("%d ", prime)
	}
}
