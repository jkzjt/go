package main

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Feet  float64
	Meter float64
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return
	}
	feet := Feet(f)
	meter := Meter(feet) * .3048
	fmt.Printf("%g Feet is %.2f Meters\n", feet, meter)
}
