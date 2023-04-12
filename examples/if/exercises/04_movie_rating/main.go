package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	tip     = "Requires age"
	errAge  = "Wrong age: %q\n"
	errArg  = "Wrong arg: %q\n"
	ab17    = "R-Rated"
	bt13a17 = "PG-13"
	bl13    = "PG-Rated"
)

const r, pg = 17, 13

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}

	if age, err := strconv.Atoi(strings.TrimSpace(os.Args[1])); err == nil {
		if age < 0 {
			fmt.Printf(errAge, os.Args[1])
			return
		}
		if age > r {
			fmt.Println(ab17)
		} else if age >= pg {
			fmt.Println(bt13a17)
		} else {
			fmt.Println(bl13)
		}
	} else {
		fmt.Printf(errArg, os.Args[1])
	}

}
