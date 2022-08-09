package main

import "fmt"

func closure() func() {
	x := 10

	funcao := func() {
		fmt.Println(x)
	}

	return funcao
}

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number. Typed number: %d", n)
	case n == 0:
		return 1, nil
	default:
		previousFactorial, _ := factorial(n - 1)
		return n * previousFactorial, nil
	}
}

func main() {
	x := 20

	fmt.Println(x) // 20

	printX := closure()

	printX() // 10

	// Factorial

	result, _ := factorial(5)

	fmt.Println(result)

	_, err := factorial(-4)

	if err != nil {
		fmt.Println(err)
	}
}
