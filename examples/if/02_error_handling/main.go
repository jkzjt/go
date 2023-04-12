package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Itoa returns no errors
	// Don't need to handle for it
	// func Itoa(i int) string
	fmt.Println(strconv.Itoa(10))

	// Atoi returns an error that should
	// be checked and handled if needed
	// func Atoi(s string) (int, error)
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
