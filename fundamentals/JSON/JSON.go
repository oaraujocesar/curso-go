package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{ID: 1, Name: "Ring", Price: 599.30, Tags: []string{"Wearable", "Jewel", "Gold"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	//json to struct
	var p2 product

	jsonString := `{"id": 2, "name": "Pen", "price": 8.9, "tags": ["Paper", "Tool", "Drawing", "Writing"]}`

	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
