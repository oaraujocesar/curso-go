package main

import (
	"fmt"

	"github.com/oaraujocesar/html"
)

func forward(origin <-chan string, target chan string) {
	for {
		target <- <-origin
	}
}

// multiplexer - mixup (messages) in one channel
func group(sources ...<-chan string) <-chan string {
	c := make(chan string)

	for _, source := range sources {
		go forward(source, c)
	}

	return c
}

func main() {
	c := group(
		html.Title("https://www.cod3r.com.br", "https://www.google.com"),
		html.Title("https://www.amazon.com", "https://youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
