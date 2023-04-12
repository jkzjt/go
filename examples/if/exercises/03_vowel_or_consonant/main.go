package main

import (
	"fmt"
	"os"
)

const a, e, i, o, u, y, w = "a", "e", "i", "o", "u", "y", "w"

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
	if lt == "a" || lt == "e" || lt == "i" || lt == "o" || lt == "u" {
		fmt.Printf(vow, lt)
	} else if lt == "y" || lt == "w" {
		fmt.Printf(smv, lt)
	} else {
		fmt.Printf(csn, lt)
	}
}
