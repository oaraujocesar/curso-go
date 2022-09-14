package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// speak("César", "Talk to me!", 3)
	// speak("John", "After you", 1)

	// go speak("César", "Talk to me!", 500)
	// go speak("John", "After you", 500)

	// time.Sleep(time.Second * 30)

	// fmt.Println("Fim")

	go speak("César", "Talk to me!", 10)
	speak("Fran", "Congrats!", 4)
}
