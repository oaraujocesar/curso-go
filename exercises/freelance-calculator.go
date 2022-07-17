package main

import (
	"fmt"

	c "golang.org/x/text/currency"
	l "golang.org/x/text/language"
	m "golang.org/x/text/message"
	n "golang.org/x/text/number"
)

func main() {
	var monthlyPay float64
	var dailyHours int64

	fmt.Println("How much do you make in a month?")
	fmt.Scanf("%f", &monthlyPay)

	fmt.Println("What's your daily working period?")
	fmt.Scanf("%d", &dailyHours)

	valuePerHour := monthlyPay / float64(dailyHours*22)

	lang := l.MustParse("en_US")
	cur, _ := c.FromTag(lang)

	scale, _ := c.Cash.Rounding(cur)
	dec := n.Decimal(valuePerHour, n.Scale(scale))

	p := m.NewPrinter(lang)
	p.Printf("You gonna earn: %v%v per hour", c.Symbol(cur), dec)
}
