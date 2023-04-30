package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// How to get all the valid runes in a string 
	s := "Hello, 世界. Olá, ゴーファー."
	
	// %q + []rune(...)
	// fmt.Println([]rune(s)) // NO
	fmt.Printf("[]rune   : %q\n", []rune(s))
	
	fmt.Println()
	// %q + for range
	fmt.Print("for range: [")
	for _, r := range s {
		fmt.Printf("%q ", r)
	}
	fmt.Print("]")
	
	fmt.Println()
	fmt.Println()
	
	// %q + for + len() + DecodeRuneInString()
	// https://pkg.go.dev/unicode/utf8@latest#DecodeRuneInString
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if i == 0 {
			fmt.Printf("for      : [%q ", r)
		} else if i + size == len(s) {
			fmt.Printf("%q]", r)
		}else {
			fmt.Printf("%q ", r)
		}
		
		i += size
	}
	
	fmt.Println()
	fmt.Println()
	
	// How to get all the valid runes in a byte slice
	// See https://pkg.go.dev/unicode/utf8@latest#pkg-functions
}