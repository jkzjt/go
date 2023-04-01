package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	feet, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid argument: %q\n", os.Args[1])
		return
	}
	meter := feet * .3048
	fmt.Printf("%g feet is %.2f meters.\n", feet, meter)
}
