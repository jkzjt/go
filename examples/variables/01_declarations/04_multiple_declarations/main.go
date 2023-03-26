package main

import (
	"fmt"
)

func main() {
	// -----------------------------------------------------
	// 3 forms in variable declaration are equivalent below
	// -----------------------------------------------------

	// separate
	// var title string
	// var stock int
	// var price float64
	// var onSale bool

	// multiple
	var (
		title  string
		stock  int
		price  float64
		onSale bool
	)

	// parallel -- declaring variables in this way the variables must be same data type
	// here `any` is an alias for `interface{}`, an empty interface
	// var title, stock, price, onSale any

	fmt.Printf(
		"{\n\ttitle: %q\n\tstock: %d\n\tprice: %.2f\n\tonSale: %t\n}",
		title, stock, price, onSale,
	)
}
