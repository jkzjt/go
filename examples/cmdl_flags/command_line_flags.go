package main

import (
	"flag"
	"fmt"
)

func main() {
	strPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("num", 42, "a number")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *strPtr)
	fmt.Println("num: ", *numPtr)
	fmt.Println("bool: ", *boolPtr)
	fmt.Println("svar: ", svar)
	fmt.Println("tail: ", flag.Args())
}
