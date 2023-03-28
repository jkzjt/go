package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	formatter := "My name is %s and my surname is %s.\n"
	fmt.Printf(formatter, os.Args[1], os.Args[2])
}
