package main

import (
	"fmt"
)

// NOTE THAT: CANNOT USE SHORT DECLARATION AT THE PACKAGE SCOPE
// following produces a compile error
// SYNTAX ERROR: non-declaration statement outside function body
// ok := true

// sure, `var` keyword can be used for variable declaration at
// the package scope, like below

// #1 have no idea on initializer value, it would use zero value
// here is `false`
// var ok bool

// #2 need the specific value
// var ok bool = true
var ok = true

func main() {
	fmt.Println(ok)
}
