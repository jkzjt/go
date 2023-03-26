package main

import (
	"fmt"
)

func main() {
	base := 100
	factor := 2.5

	// BUILD FAIL
	// invalid operation: base * factor (mismatched types int and float64)
	// base := base * factor

	// conversion is a dangerous operation
	// may yield a loss of precision
	// here `factor` loses it's fractional part
	// base is 200, it's destructive
	// base = base * int(factor)

	// so conversion order matters
	base = int(float64(base) * factor)

	fmt.Println(base)
}
