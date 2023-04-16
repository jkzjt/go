package main

import . "fmt"

func main() {
	var (
		sum int
		i = 1
	)
	// keyword `goto`
loop:
	if i <= 5 {
		sum += i
		i++
		goto loop
	}
	Println(sum) // 15
	
	// i == 1 -> i <= 5(true)  -> sum += i(1)  -> i++(2)
	// i == 2 -> i <= 5(true)  -> sum += i(3)  -> i++(3)
	// i == 3 -> i <= 5(true)  -> sum += i(6)  -> i++(4)
	// i == 4 -> i <= 5(true)  -> sum += i(10) -> i++(5)
	// i == 5 -> i <= 5(true)  -> sum += i(15) -> i++(6)
	// i == 6 -> i <= 5(false) X
}