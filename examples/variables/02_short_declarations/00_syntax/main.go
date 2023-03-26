package main

import (
	"fmt"
)

func main() {
	// `=` the assignment operator
	// when used with variable declaration, it
	// initializes the variable to the given value

	// #1 below initializes `safe` to true
	// var safe bool = true

	// #2
	// -- var variable_name = initializer value
	// GO infers `safe`'s type from the initializer value
	// here `true`'s type is bool, so `safe`'s type becomes bool
	// var safe = true

	// #3 short declaration -- WITHOUT `var` KEYWORD
	// -------------------------------------------
	// SYNTAX:
	// 			variable_name := initializer value
	// -------------------------------------------
	safe := true

	// #1,#2,#3 are equivalent

	fmt.Printf("safe(%T)? %[1]t\n", safe)

	// type inference mechanism also works for variables in GO
	// following is an instance that GO gets the type
	// of `on` and assigns it to the `activated`
	// the type of both are same `bool`
	on := true
	activated := on

	fmt.Printf(
		"on(%T)?%[1]t\nactivated(%T)?%[2]t\n",
		on, activated,
	)

}
