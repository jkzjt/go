package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	lower  = 6
	upper  = 16
	lenErr = "The length of nickname must be 6~16"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	nickname := strings.TrimSpace(os.Args[1])
	l := utf8.RuneCountInString(nickname)
	if l < lower || l > upper {
		fmt.Println(lenErr)
		return
	}
	fmt.Printf("%q\n", os.Args[1])
	fmt.Println(nickname)
}
