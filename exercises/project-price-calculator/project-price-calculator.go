package main

import (
	"fmt"

	c "golang.org/x/text/currency"
	l "golang.org/x/text/language"
	m "golang.org/x/text/message"
	n "golang.org/x/text/number"
)

func formatToCurrency(value float64) (n.Formatter, c.Unit, *m.Printer) {
	lang := l.MustParse("en_US")
	cur, _ := c.FromTag(lang)

	scale, _ := c.Cash.Rounding(cur)
	dec := n.Decimal(value, n.Scale(scale))

	p := m.NewPrinter(lang)

	return dec, cur, p
}

func main() {

	var rate float64
	var hours int64

	fmt.Println("What's your hourly rate?")
	fmt.Scanf("%f", &rate)

	fmt.Println("How many hours this projects is gonna take?")
	fmt.Scanf("%d", &hours)

	totalPrice := rate * float64(hours)

	dec, cur, printer := formatToCurrency(totalPrice)

	printer.Printf("Your project is gonna cost: %v%v", c.Symbol(cur), dec)

}
