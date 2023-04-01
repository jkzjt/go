package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	r := []rune(os.Args[1])
	l := len(r)
	fmt.Println(
		os.Args[1] + strings.Repeat(string(r[l-1]), l),
	)
}
