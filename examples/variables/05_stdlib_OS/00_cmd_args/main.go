package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Args' length: %d \n", len(os.Args))
	
	for i, e := range os.Args {
		fmt.Printf("Args[%d] = %s\n", i, e)
	}
}