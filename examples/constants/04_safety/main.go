package main

func main() {
	// ---------------------------------------------------
	// constants' initializer values are determined at the
	// compile-time, and errors on which are checked
	// ---------------------------------------------------

	// panic: runtime error: integer divide by zero
	n, m := 1, 0
	// Println(n / m)

	// build fail: invalid operation: division by zero
	// Println(1 / 0)

	// build fail: invalid operation: division by zero
	const a, b = 1, 0
	// Println(a / b)

	// build fail: invalid operation: division by zero
	// Println(n / 0)
	// panic: runtime error: integer divide by zero
	// Println(1 / m)
	// build fail: invalid operation: division by zero
	// Println(n / b)
	// panic: runtime error: integer divide by zero
	// Println(a / m)

	// build fail: invalid operation: division by zero
	// Println(a / 0)
	// build fail: invalid operation: division by zero
	// Println(1 / b)

	_, _, _, _ = n, m, a, b
}
