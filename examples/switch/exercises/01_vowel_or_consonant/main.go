package main

import (
	"fmt"
	"os"
	"strings"
)

// --------------------------------------------------------------
// https://github.com/jkzjt/go/blob/master/examples/if/exercises/03_vowel_or_consonant/README.txt
// -------------------------------------------------

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
	lt := strings.TrimSpace(os.Args[1])
	switch lt {
	case "a", "e", "i", "o", "u":
		fmt.Printf(vow, lt)
	case "y", "w":
		fmt.Printf(smv, lt)
	default:
		fmt.Printf(csn, lt)
	}
}
