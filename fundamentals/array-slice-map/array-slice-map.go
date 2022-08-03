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

	// MAKE

	s := make([]int, 10)

	s[9] = 12

	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1)
	fmt.Println(s, len(s), cap(s))

	arr := make([]int, 10, 20)
	arr2 := append(arr, 1, 2, 3)

	fmt.Println(arr, arr2)

	arr2[3] = 12
	fmt.Println(arr, arr2)

	// COPY

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	fmt.Println(slice2)

	copy(slice2, slice1)
	fmt.Println(slice2)

}
