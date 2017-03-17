package main

import (
	"encoding/json"
	"fmt"
)

var jsonData = []byte(`[
	{
		"name": "Tim",
		"age": 75
	},
	{
		"name": "Jessica",
		"age": 42
	}
]`)

type Person struct{
	Name string `json:"name"`
	Age	int `json:"age"`
}

func main() {
	// unmarshal here
	var person []Person
	json.Unmarshal(jsonData, &person)

	var sum = 0
	for i := range person{
		sum += person[i].Age
	}
	fmt.Printf("Sum of ages = %v", sum)
	fmt.Println()

}
