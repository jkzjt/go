package main

import "fmt"

func main() {
	var sum float64

	// after the loop below
	// inaccuracy is clear
	for range [...]int{10: 0} {
		sum += 1.0 / 10.0
	}

	fmt.Printf("%.60f\n", sum)
}
