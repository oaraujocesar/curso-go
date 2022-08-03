package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println(strings.Repeat("=", 30))

	// SLICES
	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	for _, n := range s1 {
		fmt.Println(n * 3)
	}

	a2 := [5]uint{1, 2, 3, 4, 5}

	s2 := a2[1:3]

	fmt.Println(s2)

	fmt.Println(strings.Repeat("=", 30))

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

	fmt.Println(strings.Repeat("=", 30))

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

	fmt.Println(strings.Repeat("=", 30))

	// MAP
	// var approved map[int]string
	// maps should be initialized

	approved := make(map[int]string)

	approved[12333214565] = "Maria"
	approved[12363474573] = "Pedro"
	approved[34959923451] = "Anna"

	fmt.Println(approved)
	fmt.Println(approved[12333214565]) // read
	delete(approved, 12333214565)
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	salaries := map[string]float64{
		"César O. Araújo": 32250.30,
		"Chava The Golem": 35930.11,
		"The Djim":        6753.56,
	}

	salaries["Rafael Filho"] = 1350.0

	delete(salaries, "nonexistent")

	for name, salary := range salaries {
		fmt.Println(name, salary)
	}

	nested := map[string]map[string]float64{
		"C": {
			"César": 7000.0,
			"Clara": 8520.33,
		},
		"D": {
			"Daniel": 12344.10,
		},
	}

	delete(nested, "D")

	fmt.Println(strings.Repeat("=", 30))

	for letter, employees := range nested {
		fmt.Println(letter, employees)
	}
}
