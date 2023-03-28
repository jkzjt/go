package main

import "fmt"

func main() {
	var (
		n int = -1
		un uint = 1
		fl float64 = 1.0
		bl bool = true
		str string = "HELLO"
	)
	
	// `%v`'s default format is `%d`
	fmt.Printf("%v\n", n)
	
	// `%v`'s default format is `%d`
	// `%x` if printed with `%#v`
	fmt.Printf("%v\n", un)
	fmt.Printf("%#v\n", un)
	
	// `%v`'s default format is `%g`
	fmt.Printf("%v\n", fl)
	
	// `%v`'s default format is `%t`
	fmt.Printf("%v\n", bl)
	
	// `%v`'s default format is `%s`
	fmt.Printf("%v\n", str)
	
	// `%T` -- a Go syntax representation of the type of the value
	// int
	fmt.Printf("%T\n", n)
	// uint
	fmt.Printf("%T\n", un)
	// float64
	fmt.Printf("%T\n", fl)
	// bool
	fmt.Printf("%T\n", bl)
	// string
	fmt.Printf("%T\n", str)
	
	// %% -- a literal percent sign; consumes no value
	fmt.Printf("%%\n")
	
	// %#v a Go syntax representation of the value
	fmt.Printf("%#v\n", n)
	fmt.Printf("%#v\n", un)
	fmt.Printf("%#v\n", fl)
	fmt.Printf("%#v\n", bl)
	fmt.Printf("%#v\n", str)
}