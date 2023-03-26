package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(a, &a)

	a, b := 2, 10
	fmt.Println(a, &a, b, &b)

	if a, b := 3, 11; true {
		fmt.Println(a, &a, b, &b)
		a := 4
		fmt.Println(a, &a, b, &b)
		b := 12
		fmt.Println(a, &a, b, &b)
		
		// BUILD FAIL
		// no new variables on left side of :=
		// a, b := 5, 13
	}
	
	fmt.Println(a, &a, b, &b)
}
