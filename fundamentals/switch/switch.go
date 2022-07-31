package main

import (
	"fmt"
	"time"
)

func gradeToConcept(g float64) string {
	var grade = int(g)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade."
	}
}

func whichType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknown"
	}
}

func main() {
	concept := gradeToConcept(10)
	fmt.Println(concept)
	fmt.Println("==========================")

	concept2 := gradeToConcept(5)
	fmt.Println(concept2)
	fmt.Println("==========================")

	concept3 := gradeToConcept(1)
	fmt.Println(concept3)
	fmt.Println("==========================")

	concept4 := gradeToConcept(123)
	fmt.Println(concept4)
	fmt.Println("==========================")

	t := time.Now()

	switch { // same as switch true {}
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	fmt.Println(whichType(2.3))
}
