package main

import "fmt"

func main() {
	var grades [3]float64

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	fmt.Println(grades)

	var total float64

	count := float64(len(grades))

	for _, num := range grades {
		total += num
	}

	average := total / count

	fmt.Printf("%.2f", average)
}
