package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

type item struct {
	id       int
	quantity int
	price    float64
}

type request struct {
	userID int
	items  []item
}

// method: function with a receiver
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount/100)
}

func (r request) amount() float64 {
	total := 0.0

	for _, item := range r.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	product1 := product{
		name:     "Pen",
		price:    1.99,
		discount: 10,
	}

	fmt.Printf("US$: %.2f\n", product1.priceWithDiscount())

	request := request{
		userID: 1,
		items: []item{
			{id: 1, quantity: 2, price: 12.10},
			{id: 2, quantity: 5, price: 7.99},
			{id: 11, quantity: 100, price: 3.13},
		},
	}

	fmt.Printf("The total amount of the request is US$ %.2f", request.amount())
}
