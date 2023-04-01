package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// `Hello World!` in Ukrainian
	// it contains unicode characters
	str := "Привіт, світе!"

	// 25 not 14
	fmt.Println(len(str))

	// To get the actual characters(runes) inside it
	// we should use function `RuneCountInString` under
	// package `unicode/utf8`
	// 14 right
	fmt.Println(utf8.RuneCountInString(str))
}
