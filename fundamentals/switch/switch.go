package main

import "fmt"

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
}
