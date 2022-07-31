package main

import "fmt"

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

func main() {
	approveOrNot(6)
	approveOrNot(7.3)

	result := approveOrNotWithConcept(7)
	fmt.Println(result)
}
