package main

import . "fmt"

func main() {
	{
		if /*#1*/ a := 1; true {
			Println(&a, a)
			/*#2*/ a := a * 10
			Println(&a, a)
		}

		// -----------------------------------------------------------
		// `&` is a adress operator
		// #1 and #2 -- `a` are different variables in different scopes
		// -------------------------------------------------------
	}

	{
		if /*#1*/ a := 1; false {
			//
		} else if /*#2*/ a := &a; true {
			Printf("%p %p %d\n", &a, a, *a) // ? ? 1
		}

		// -----------------------------------------------------------
		// `*` in `*a` is a dereference operator
		// #1 and #2 -- `a` are different variables in different scopes
		// -------------------------------------------------------
	}

	{
		if a := 1; false { // int, a = 1
			//
		} else if a := &a; false { // *int, *a = 1
			//
		} else if a := &a; true { // **int, **a = 1
			// &a -- ***int
			Printf("%p %p %p %d\n", &a, a, *a, **a) // ? ? ? 1
		}
		// ...
	}
	
	{
		
		// var (
		//	a int
		//	b = &a
		// )
		// build fail: invalid operation: cannot indirect *b (variable of type int)
		// Printf("%d\n", **b)
	}
}
