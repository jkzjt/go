package main

import (
	"fmt"
)

func main() {

	// numeric type
	var age int        // zero value is 0
	var weight float64 // zero value is 0.0

	// bool constants
	var married bool // zero value is false

	// string
	var fullname string // zero value is ""

	fmt.Printf(
		"age: %d\nweight: %.0f\nmarried: %t\nfullname: %q\n",
		age, weight, married, fullname,
	)

}
