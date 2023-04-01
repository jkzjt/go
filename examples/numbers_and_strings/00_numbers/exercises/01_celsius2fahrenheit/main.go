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
	c, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid argument: %q\n", os.Args[1])
		return
	}
	f := c*1.8 + 32.
	fmt.Printf("%g ºC is %.2f ºF.\n", c, f)
}
