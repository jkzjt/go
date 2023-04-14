package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ----------------------------------------------------------
// https://github.com/jkzjt/go/blob/master/examples/if/exercises/07_days_in_month/README.txt
// -------------------------------------------------------------

const (
	tip    = "Give me a month"
	msg    = "%q has %d days.\n"
	errMsg = "%q is not a month.\n"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	s := strings.TrimSpace(os.Args[1])
	m := strings.ToLower(s)
	switch m {
	case "january", "march", "may", "july", "august", "october", "december":
		fmt.Printf(msg, s, 31)
	case "april", "june", "september", "november":
		fmt.Printf(msg, s, 30)
	case "february":
		y := time.Now().Year()
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			fmt.Printf(msg, s, 29)
		} else {
			fmt.Printf(msg, s, 28)
		}
	default:
		fmt.Printf(errMsg, s)
	}
}
