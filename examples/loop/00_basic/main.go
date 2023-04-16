package main 

import . "fmt"

func main() {
	var sum int
	
	// keyword `for`
	for i := 1; i <= 5; i++ {
		sum += i
	}
	Println(sum) // 15
	
	// i == 1 -> i <= 5(true)  -> sum += i(1)  -> i++(2)
	// i == 2 -> i <= 5(true)  -> sum += i(3)  -> i++(3)
	// i == 3 -> i <= 5(true)  -> sum += i(6)  -> i++(4)
	// i == 4 -> i <= 5(true)  -> sum += i(10) -> i++(5)
	// i == 5 -> i <= 5(true)  -> sum += i(15) -> i++(6)
	// i == 6 -> i <= 5(false) X
	
	// while ???
	// below code is equivalent to `while true {...}`
	// uncomment it and run that will yield timeout
	
	// i := 1
	// for {
	//	  sum += i
	//	  i++
	// }
	// Println(sum, i) unreachable
}