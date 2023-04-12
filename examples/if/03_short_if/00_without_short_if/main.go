package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println(i) // 0
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	fmt.Println(err)
}
