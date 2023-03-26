package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	visitors := os.Args[1:]
	fmt.Printf("There are %d visitors.\n", len(visitors))
	for _, e := range visitors {
		fmt.Println("Welcome,", e)
	}
	fmt.Println("Welcome and enjoy your time, visitors")
}
