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
	v := 4 * math.Pi * math.Pow(r, 3) / 3
	fmt.Printf(
		"The volume of sphere with radius %g is %.2f.\n",
		r, v,
	)
}
