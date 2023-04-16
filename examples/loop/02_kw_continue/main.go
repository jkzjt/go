package main 

import . "fmt"

func main() {
	var sum int
	
	// keyword `continue`
	for i := 1; i <= 5; i++ {
		// skip the current loop portion when meet
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	Println(sum) // 9
	
	// i == 1 -> i <= 5(true)  -> i%2 == 0(false) -> sum += i(1) -> i++(2)
	// i == 2 -> i <= 5(true)  -> i%2 == 0(true)  -> i++(3)
	// i == 3 -> i <= 5(true)  -> i%2 == 0(false) -> sum += i(4) -> i++(4)
	// i == 4 -> i <= 5(true)  -> i%2 == 0(true)  -> i++(5)
	// i == 5 -> i <= 5(true)  -> i%2 == 0(false) -> sum += i(9) -> i++(6)
	// i == 6 -> i <= 5(false) X
}