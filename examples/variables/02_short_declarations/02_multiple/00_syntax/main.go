package main

import (
	"fmt"
)

func main() {
	// #1
	// fullname := "ilauyu"
	// age := 18
	// weight := 49.9
	// married  := true

	// #2 is equivalent to #1
	// the number of variables and values must be equal
	// variables are assigned in order
	// it supports differrent data types
	// unlimited on the numder of variables and values 
	// but declare more, read worse
	fullname, age, weight, married := "ilauyu", 18, 49.9, true

	fmt.Printf(
		"{\n\tfullname: %s\n\tage: %d\n\tweight: %.1f\n\tmarried: %t\n}",
		fullname, age, weight, married,
	)

}
