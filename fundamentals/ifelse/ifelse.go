package main

import "fmt"

func approveOrNot(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade:", grade)
	} else {
		fmt.Println("Disapproved with grade:", grade)
	}
}

func main() {
	approveOrNot(6)
	approveOrNot(7.3)
}
