package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	fmt.Printf(
		" IN: %q\nOUT: %q\n",
		os.Args[1], strings.TrimRight(os.Args[1], " "),
	)
}
