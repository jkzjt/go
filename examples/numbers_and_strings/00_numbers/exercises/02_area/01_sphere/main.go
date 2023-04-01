package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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
	area := 4 * math.Pi * math.Pow(r, 2)
	fmt.Printf(
		"The area of sphere with radius %g is %.2f.\n",
		r, area,
	)
}
