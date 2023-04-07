package main

import (
	"fmt"
	"time"
)

// with predeclared types, why we need
// to declare our types?
// sometimes for distinction, readability and type-safety
// sometimes for actions

// type definitions usually declared at package scope
type (
	Meter     float64
	Kilometer float64
)

// TODO
// define methods on type `meter` `kilometer`

func main() {
	// look at the example first from stdlib
	{
		// fmt.Println(time.Now())
		h, err := time.ParseDuration("14h45m")
		if err != nil {
			return
		}
		fmt.Println(h.Hours())
		var m int64 = 2

		// invalid op
		// mismatched types
		// h *= m

		// conversion is needed
		h *= time.Duration(m)
		fmt.Println(h)

		// type Duration int64
		fmt.Printf("type of `h`: %T\n", h)
		fmt.Printf("underlying type of `h`: %T\n", int64(h))
	}

	{
		var (
			m  Meter = 100
			km Kilometer
		)

		// meter and kilometer are different types
		// but the underlying of both are same, float64
		// so they are convertable to each other

		// BUILD FAIL
		// type mismatched
		// km = m / 1000.

		// works
		km = Kilometer(m) / 1000.
		fmt.Printf("%g m is %g km.\n", m, km)

		// underlying type
		fmt.Printf(
			"underlying type of `m` %T\nunderlying type of `km` %T\n",
			float64(m), float64(km),
		)

		fmt.Printf("%T\n", m)
		fmt.Printf("%T\n", km)
	}

	{
		type (
			Centimeter float64
			Decimeter  Centimeter
		)

		var (
			cm Centimeter = 1000
			dm Decimeter
		)
		// BUILD FAIL
		// dm = cm

		fmt.Printf("%T\n", cm)
		fmt.Printf("%T\n", dm)
	}
}
