package main

import (
	"fmt"
	"strconv"
)

func main() {
	// `%b` verb prints minimum bits which present the given value
	// with no leading zeros
	
	// 1 bit can encode 2 different states
	// 0 or 1
	fmt.Printf("%b\n", 0)
	fmt.Printf("%b\n", 1)
	
	// `%0nb` verb prints n bits provied representing the given value
	// with leading zeros
	
	// 2 bit can encode 4 different states
	// 00, 01, 10 or 11
	fmt.Printf("%02b = %[1]d\n", 0)
	fmt.Printf("%02b = %[1]d\n", 1)
	fmt.Printf("%02b = %[1]d\n", 2)
	fmt.Printf("%02b = %[1]d\n", 3)
	
	// n bit can encode 2^n different states
	
	i, err := strconv.ParseInt("00001000", 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%08b = %[1]d\n", i)
	
	i, err = strconv.ParseInt("01011010", 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%08b = %[1]d\n", i)
	
}