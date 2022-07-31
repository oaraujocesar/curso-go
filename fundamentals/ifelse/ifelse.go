package main

import (
	"fmt"
	"math/rand"
	"time"
)

func approveOrNot(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade:", grade)
	} else {
		fmt.Println("Disapproved with grade:", grade)
	}
}

func approveOrNotWithConcept(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 7 && grade < 9 {
		return "B"
	} else {
		return "C"
	}
}

func randomNum() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	approveOrNot(6)
	approveOrNot(7.3)

	result := approveOrNotWithConcept(7)
	fmt.Println(result)

	if i := randomNum(); i > 5 {
		fmt.Println("You won")
	} else {
		fmt.Println("You lost")
	}
}
