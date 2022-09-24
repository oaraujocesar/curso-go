package mathmagic

import (
	"fmt"
	"strconv"
)

/*
	Average returns the average (Aha!)

it takes infinite float64 numbers and returns the average
*/
func Average(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	avg := total / float64(len(numbers))
	roundedAvg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)

	return roundedAvg
}
