package main

import (
	"fmt"
	"os"
)

func main() {
	l := len(os.Args)
	if l == 1 {
		fmt.Println("Give me args")
		return
	}
	if l == 2 {
		fmt.Printf("There is one: %q\n", os.Args[1])
	} else if l == 3 {
		// fmt.Printf("There are two: %q\n", os.Args[1]+" "+os.Args[2])
		fmt.Printf(`There are two: "%s %s"`+"\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguments", l-1)
	}
}
