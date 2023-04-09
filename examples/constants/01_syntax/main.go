package main

import . "fmt"

func main() {
	// -------------------------------------------
	// keyword `const`
	// SYNTAX
	//  const identifier [type] = initializer value
	// --------------------------------------------

	// const meter int = 10 // here meter is a named and typed constant

	// no need to specify types, the compiler will infer for us
	// --------------------------------------------
	// 100  	     -> int
	// 100. 	     -> float64
	// true, false   -> bool
	// `Hi`, "Hello" -> string
	// '谭'， 'A'    -> rune(an alias for int32)
	// --------------------------------------------
	const meter = 10 // here meter is a named constant
	Printf("%T\n", meter)

	// a named constant must be declared and assigned at the
	// meanwhile whether it is typed or not
	// the value of a constant is determined at compile-time
	// const meter int // BUILD FAIL: missing init expr for meter
	// const meter // BUILD FAIL: missing init expr for meter

	// constants in GO have their own zero values?
	// yes, but we need to manually assign a initializer value

	dm := 10
	m := dm / meter
	_, _ = dm, m

	// -------------------------------------------------------------
	// parallel declaration
	// SYNTAX
	// const identifier0, identifier1, ... [type] = value0, value1, ...
	// ----------------------------------------------------------------
	// const a, b, c int = 10, 20, 30
	const a, b, c = 10, 20, 30
	_, _, _ = a, b, c

	// ---------------------------------------
	// group declaration
	// SYNTAX
	// const (
	// 		identifier0 [type] = value0
	//		identifier1 [type] = value1
	//		...         ...    ...
	// )
	// ---------------------------------------
	// const (
	//	e int     = 1
	//	f float64 = 1.2
	// )
	const (
		e = 1
		f = 1.2
	)
	_, _ = e, f
}
