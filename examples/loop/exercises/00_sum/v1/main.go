package main

import . "fmt"

// ------------------------------------
// sum the odd numbers between 1 and 100
// 
// EXPECTED OUTPUT
//  2500
// --------------------------------------

func main() {
	var sum int
	
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i;
	}
	Println(sum)
}