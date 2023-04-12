package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	vowel   = "aeiou"
	semivow = "yw"
)

const (
	tip = "Give me a letter."
	vow = "%q is a vowel.\n"
	csn = "%q is a consonant.\n"
	smv = "%q is sometimes a vowel, sometimes not.\n"
)

func main() {
	if len(os.Args) == 1 || len(os.Args[1]) != 1 {
		fmt.Println(tip)
		return
	}
	lt := os.Args[1]
	
	// ------------------------
	// Use strings#ContainsAny()
	// ------------------------
	
	if strings.ContainsAny(vowel, lt) {
		fmt.Printf(vow, lt)
	} else if strings.ContainsAny(semivow, lt) {
		fmt.Printf(smv, lt)
	} else {
		fmt.Printf(csn, lt)
	}
}