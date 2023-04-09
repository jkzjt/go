package main

import "fmt"

func main() {
	// -------------------------------------
	// The initializer values of constants
	// can be an  constant expression
	// --------------------------------------
	{
		const (
			sum  = 1 + 1
			on   = !false
			ver  = `2.0.1` + "-beta"
			rate = .3 + .5
		)
		_, _, _, _ = sum, on, ver, rate
	}
	// ----------------------------------------------
	// If type is omitted with no assignment statement
	// when using group declaration of constant, the
	// constant will repeat the previous type and expression
	// provided any
	// -----------------------------------------
	{
		const (
			a = 97
			b
			m = .0 + 56
			n
		)
		fmt.Println(a, b, m, n)                 // 97 97 56 56
		fmt.Printf("%T %T %T %T\n", a, b, m, n) // int int float64 float64
	}

	{
		const (
			Sunday    = 0
			Monday    = 1
			Tuesday   = 2
			Wednesday = 3
			Thursday  = 4
			Friday    = 5
			Saturday  = 6
		)
		fmt.Println(
			Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday,
		)
	}

	// -----------------------------------------------------------------
	// iota is a predeclared identifier representing the untyped integer ordinal
	// number of the current const specification in a (usually parenthesized)
	// const declaration. It is zero-indexed.
	// --------------------------------------------
	{
		// build fail: cannot use iota outside constant declaration
		// fmt.Println(iota)
		{
			const (
				a = iota // 0
				b        // 1
				c        // 2
			)
			fmt.Println(a, b, c)
		}
		{
			const (
				a = iota // 0
				b        // 1
				_        // 2
				c        // 3
			)
			fmt.Println(a, b, c)
		}
		{
			const (
				a = iota // 0
				b        // 1

				c // 2
			)
			fmt.Println(a, b, c)
		}
		{
			const (
				a = iota // 0
				b        // 1
				c = true
				d        // true
				e = iota // 4
			)
			fmt.Println(a, b, c, d, e)
		}
		{
			const (
				a = 2
				b = 'a'
				c
				d = `Hi`
				e = iota // 4
			)
			fmt.Println(a, b, c, d, e) // 2 97 97 Hi 4
		}
		{
			const (
				a = 2
				b = 'a'
				c = iota // 2
			)
			const g = iota // 0 reset
			fmt.Println(a, b, c, g)
		}
	}

	{
		const (
			Sunday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)
		fmt.Println(
			Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday,
		)
	}
}
