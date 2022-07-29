package main

import "fmt"

func purchases(firstJob, secondJob bool) (bool, bool, bool) {
	buy50InchTv := firstJob && secondJob
	buy32InchTv := firstJob != secondJob // exclusive or
	buyIceCream := firstJob || secondJob

	return buy50InchTv, buy32InchTv, buyIceCream
}

func main() {
	tv50, tv32, iceCream := purchases(true, true)

	fmt.Println(tv50, tv32, iceCream)
}
