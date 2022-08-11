package main

import "fmt"

type grade float64

func (g grade) inbetweens(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeToConcept(g grade) (concept string) {

	switch {
	case g.inbetweens(9.0, 10.0):
		return "A"
	case g.inbetweens(7.0, 8.99):
		return "B"
	case g.inbetweens(5.0, 7.99):
		return "C"
	case g.inbetweens(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	var i grade = 3

	fmt.Println(gradeToConcept(i))
}
