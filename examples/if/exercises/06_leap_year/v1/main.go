package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const m4, m100, m400 = 4, 100, 400

const (
	tip    = "Give me a year number"
	errMsg = "%q is an invalid year number\n"
	ly     = "%d is a leap year\n"
	nly    = "%d is not a leap year\n"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	s := strings.TrimSpace(os.Args[1])
	y, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf(errMsg, s)
		return
	}

	var leap bool
	if y%m400 == 0 {
		leap = true
	} else if y%m100 == 0 {
		leap = false
	} else if y%m4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf(ly, y)
	} else {
		fmt.Printf(nly, y)
	}
}
