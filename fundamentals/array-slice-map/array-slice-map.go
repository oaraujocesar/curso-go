package main

import "fmt"

func main() {
	// ARRAYS
	var grades [3]float64

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	fmt.Println(grades)

	var total float64

	count := float64(len(grades))

	for _, num := range grades {
		total += num
	}

	average := total / count

	fmt.Printf("%.2f\n", average)

	numbers := [...]int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i, number)
	}

	// SLICES
	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	for _, n := range s1 {
		fmt.Println(n * 3)
	}

	a2 := [5]uint{1, 2, 3, 4, 5}

	s2 := a2[1:3]

	fmt.Println(s2)

}
