package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	tip = "Type magnitude in human terms"
	msg = "%s's richter scale is %s\n"
	mas = "10+"
	gre = "8 - 9.9"
	maj = "7 - 7.9"
	str = "6 - 6.9"
	mod = "5 - 5.9"
	lig = "4 - 4.9"
	min = "3 - 3.9"
	ver = "2 - 2.9"
	mic = "0.1 - 1.9"
	unk = "unknown"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	switch s := strings.TrimSpace(os.Args[1]); s {
	case "massive":
		fmt.Printf(msg, s, mas)
	case "great":
		fmt.Printf(msg, s, gre)
	case "major":
		fmt.Printf(msg, s, maj)
	case "strong":
		fmt.Printf(msg, s, str)
	case "moderate":
		fmt.Printf(msg, s, mod)
	case "light":
		fmt.Printf(msg, s, lig)
	case "minor":
		fmt.Printf(msg, s, min)
	case "very minor":
		fmt.Printf(msg, s, ver)
	case "micro":
		fmt.Printf(msg, s, mic)
	default:
		fmt.Printf(msg, s, unk)
	}
}
