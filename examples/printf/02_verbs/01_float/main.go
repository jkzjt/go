package main

import "fmt"

func main() {
	n := 15.5
	// decimalless scientific notation with exponent a power of two
	// in the manner of strconv.FormatFloat with the 'b' format, e.g. -123456p-78
	// 8725724278030336p-49
	fmt.Printf("%b\n", n)
	
	// scientific notation
	// 1.550000e+01
	fmt.Printf("%e\n", n)
	
	// scientific notation
	// 1.550000E+01
	fmt.Printf("%E\n", n)
	
	// decimal point but no exponent
	// 15.500000
	fmt.Printf("%f\n", n)
	
	// decimal point but no exponent
	// 15.500000
	fmt.Printf("%F\n", n)
	
	// %e for large exponents, %f otherwise
	// 15.5
	fmt.Printf("%g\n", n)
	
	// %E for large exponents, %F otherwise
	// 15.5
	fmt.Printf("%G\n", n)
	
	// ----------------------------
	// PRECISION
	// we generally set the precision and
	// the width is dynamic
	// if right alignment is expected, it
	// can be achieved by setting the width
	// ----------------------------
	
	num := 3.1415926
	
	// default width and default precision
	// `3.141593`
	fmt.Printf("%f\n", num)
	
	// width 9, default precision
	// ` 3.141593`
	fmt.Printf("%9f\n", num)
	
	// default width, precision 3
	// `3.142`
	fmt.Printf("%.3f\n", num)
	
	// width 9, precision 3
	// `    3.142`
	fmt.Printf("%9.3f\n", num)
	
	// width 9, precision 0
	// `        3`
	fmt.Printf("%9.f\n", num)
}