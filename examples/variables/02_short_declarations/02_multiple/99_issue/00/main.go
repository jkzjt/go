package main

import (
	"fmt"
)

// here define s struct named `cat`
type cat struct {
	// field `color`
	color string
}

func main() {

	// an instance of `cat`
	c := cat{"white"}
	fmt.Printf("%#v -- %p\n", c, &c)

	c, ok := cat{"black"}, true
	_ = ok
	fmt.Printf("%#v -- %p\n", c, &c)
	
	
	// BUILD FAIL
	// non-name c.color on left side of :=
	// c.color, on := "blue", true
	// _ = on
	// fmt.Printf("%#v -- %p\n", c, &c)
}
