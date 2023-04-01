package main

import (
	"fmt"
)

func main() {
	json := `{
	"products": [
		597426823976897931: {
			"title":  "towel", 
			"price":  19.9, 
			"amount": 2,
		}
		921749625697413668: {
			"title":  "brush", 
			"price":  24.6, 
			"amount": 3,
		}
	]
}`

	fmt.Println(json)
}