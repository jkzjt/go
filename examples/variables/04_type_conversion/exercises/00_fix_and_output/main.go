package main

import "fmt"

// --------------------------------
// fix and make output expected
// OUTPUT:
// 9.5
// -------------------------
func main() {
	// cannot convert 7.5 (untyped float constant) to type int
	// age := 2
	// fmt.Println(int(7.5) + int(age))
	
	// 7.5 (untyped float constant) truncated to int
	// age := 2
	// fmt.Println(7.5 + age)
	
	age := 2
	fmt.Println(7.5 + float64(age))
}