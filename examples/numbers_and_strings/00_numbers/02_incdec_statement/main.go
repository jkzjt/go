package main

import "fmt"

func main() {
	var n int
	// inc
	// they are equivalent
	// n = n + 1
	// n += 1
	n++
	fmt.Println(n)
	
	// dec
	// they are equivalent
	// n = n - 1
	// n -= 1
	n--
	fmt.Println(n)
	
	// WRONG
	// n = 5 + n++
	// n = -2 + n--
	// --------------------------------------
	// NOTE THAT
	// inc and dec are statements
	// --------------------------------------
	
	// they works? NO
	// ++n
	// --n
}