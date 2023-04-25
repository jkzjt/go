package main

import . "fmt"

func main() {
	// -------------------------
	// literal
	// --------------

	// var identifier [length]type = [length]type{...}
	var i [5]int = [5]int{1, 2, 3, 4, 5}

	// var identifier = [length]type{...}
	var f = [2]float64{10., 20.}

	// identifier := [length]type{...}
	a := [3]string{"World", "Gopher", "Array"}

	// ellipsis `...`
	// identifier := [...]type{...}
	// here `...` is equal to 4
	b := [...]bool{true, true, false, true}
	Printf("b: %#v\n", b)

	// mismatch type
	// build fail: cannot use [...]bool{…} (value of type [3]bool) as type [4]bool in assignment
	// b = [...]bool{true, true, true}

	_, _, _ = i, f, a
}
