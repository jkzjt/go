package main 

import . "fmt"

func main() {
	var sum int
	
	// keyword `break`
	for i := 1; ; i++ {
		// break out of the loop when meet
		if i > 5 {
			break
		}
		sum += i
	}
	Println(sum) // 15
	
	// i == 1 -> i > 5(false)  -> sum += i(1)  -> i++(2)
	// i == 2 -> i > 5(false)  -> sum += i(3)  -> i++(3)
	// i == 3 -> i > 5(false)  -> sum += i(6)  -> i++(4)
	// i == 4 -> i > 5(false)  -> sum += i(10) -> i++(5)
	// i == 5 -> i > 5(false)  -> sum += i(15) -> i++(6)
	// i == 6 -> i > 5(true) X
}