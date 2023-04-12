package main

import (
	"fmt"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi("a"); err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}
}
