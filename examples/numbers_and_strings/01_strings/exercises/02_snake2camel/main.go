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
	args := os.Args[1:]
	var b strings.Builder
	var f bool
	for _, s := range args {
		for _, r := range s {
			c := string(r)
			if c == "_" {
				f = true
				continue
			}
			if f {
				c = strings.ToUpper(c)
				f = false
			}
			b.WriteString(c)
		}
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}
