package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const factor, dividen = 2, 8

const (
	odd    = "%d is an odd number\n"
	even   = "%d is an even number\n"
	dvb8   = " and it's divisible by 8"
	tip    = "Pick a number"
	errMsg = "%s is NaN(not a number)\n"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	s := strings.TrimSpace(os.Args[1])
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf(errMsg, s)
		return
	}
	if n%factor == 0 {
		if n%dividen == 0 {
			fmt.Printf(even+dvb8, n)
		} else {
			fmt.Printf(even, n)
		}
	} else {
		fmt.Printf(odd, n)
	}
}
