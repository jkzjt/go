package main

import (
	"fmt"
)

func main() {
	// #1
	var ok bool

	fmt.Println(ok, &ok)

	// BUILD FAIL
	// no new variables on left side of :=
	// ok := true
	// ok, _ := true, 60

	// ok but make no sense
	// equivalent to `ok = true`
	// ok, _ = true, 60

	// ok
	// here `speed` is declared and given 100
	// but `ok` not, it's address is equal to #1
	// just reassigned to true
	// allowed when at least one new variable on the left of `:=`
	// NOTE THAT: `_` is not a new variable
	ok, speed := true, 100

	fmt.Println(ok, &ok)

	_ = speed

}
