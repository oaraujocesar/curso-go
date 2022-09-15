package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - readonly channel
func title(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			res, _ := http.Get(url)

			html, _ := io.ReadAll(res.Body)

			r, _ := regexp.Compile(`<title>(.*?)</title>`)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := title("https://www.cod3r.com.br/", "https://www.google.com/")
	t2 := title("https://www.amazon.com/", "https://www.youtube.com/")

	fmt.Println("Firsts: ", <-t1, "|", <-t2)
	fmt.Println("Seconds: ", <-t1, "|", <-t2)
}
