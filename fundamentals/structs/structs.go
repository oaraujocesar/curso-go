package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// method: function with a receiver
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount/100)
}

func main() {
	product1 := product{
		name:     "Pen",
		price:    1.99,
		discount: 10,
	}

	fmt.Printf("US$: %.2f", product1.priceWithDiscount())
}
