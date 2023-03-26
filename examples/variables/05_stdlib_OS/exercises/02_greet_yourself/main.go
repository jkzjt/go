package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hi, %s.\nMake your day.\n", os.Args[1])
}