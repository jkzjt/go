package main

import "fmt"

func main() {
	// multiplier (*)
	// float64
	fmt.Printf("%T\n", 5.*12)

	// division (/)
	// int
	fmt.Printf("%T\n", 9/2)

	// addition (+)
	// float64
	fmt.Printf("%T\n", 5.5+5.5)

	// subtraction (-)
	// int
	fmt.Printf("%T\n", -2-9)

	// negation (-)
	// int
	fmt.Printf("%T\n", -(-2))
	// float64
	fmt.Printf("%T\n", - -2.)

	// -------------------------------------------------
	// why are they valid above? they are literal?
	// they are constant and typeless
	// NOTE THAT
	//  when one of them is float constant, the result type
	//  is float, or int
	// -------------------------------------------------

	// remainder (%)
	// int
	fmt.Printf("%T\n", 10%3)

	// invalid op
	// `%` not defined on 10. (untyped float constant)
	// fmt.Printf("%T\n", 10.%3)

	n := 10
	// int(0)
	fmt.Printf("%T(%[1]v)\n", 100.%n)
	// `100.5` (untyped float constant) truncated to int
	// fmt.Printf("%T(%[1]v)\n", 100.5%n)
	
	// ------------------------------------------------
	// NOTE THAT
	//  `%` only works with integer
	// so why expression `100.%n` works?
	// one is `int` in it, the typeless float constant
	// would be converted to the same type without effecting
	// the original value, it's allowed
	// ------------------------------------------------
	
	// what's the value of the ratio?
	// 3 / 2 = 1.5?
	// var ratio = float64(int(3) / int(2))
	var ratio float64 = 3 / 2
	fmt.Println(ratio)
}
