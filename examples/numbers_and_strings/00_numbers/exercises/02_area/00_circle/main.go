package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	r, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid argument: %q\n", os.Args[1])
		return
	}
	// area := 3.14 * r * r
	// area := math.Pi * r * r
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf(
		"The area of a circle with radius %g is %.2f.\n",
		r, area,
	)
}
