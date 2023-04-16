package main

import . "fmt"

// ------------------------------------
// sum the odd numbers between 1 and 100 when `sum` less than 1000
// 
// EXPECTED OUTPUT
//  1024
// --------------------------------------

func main() {
	var sum int
	
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			sum += i;
		}
		if sum >= 1000 {
			break
		}
	}
	Println(sum)
}