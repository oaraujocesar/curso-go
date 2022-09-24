package main

import (
	"fmt"
	"time"

	"github.com/oaraujocesar/html"
)

func getTheFastest(url1, url2, url3 string) string {
	ch1 := html.Title(url1)
	ch2 := html.Title(url2)
	ch3 := html.Title(url3)

	select {
	case t1 := <-ch1:
		return t1
	case t2 := <-ch2:
		return t2
	case t3 := <-ch3:
		return t3
	case <-time.After(time.Second):
		return "Game over"
		// default:
		// 	return "No answer yet... :P"
	}
}

func main() {
	fastest := getTheFastest(
		"https://youtube.com",
		"https://amazon.com",
		"https://google.com",
	)

	fmt.Println(fastest)
}
