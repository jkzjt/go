package main

import . "fmt"

// ------------------------------------
// sum the numbers between 1 and 100
// 
// EXPECTED OUTPUT
//  5050
// --------------------------------------

func main() {
	var sum int
	
	for i := 1; i <= 100; i++ {
		sum += i;
	}
	Println(sum)
}