package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	{
		if i, err := strconv.Atoi("10"); err == nil {
			fmt.Println("i:", i)
			// fallthrough // fallthrough statement out of place
		} else {
			fmt.Println("err:", err)
		}

		// build fail: undefined: i, err
		// fmt.Println("i:", i)
		// fmt.Println("err:", err)
	}

	{
		if i1, err1 := strconv.Atoi("10"); err1 != nil {
			fmt.Println("err1:", err1)
		} else {
			fmt.Println("i1:", i1)
		}
	}

	{
		if args := os.Args; len(args) == 1 {
			fmt.Println("args[0]:", args[0])
		} else if i2, err2 := strconv.Atoi(args[1]); err2 == nil {
			fmt.Println("i2:", i2)
		} else {
			fmt.Println("args[1]:", args[1])
			fmt.Println("err2:", err2)
		}
	}

	{
		// #1
		var n int
		fmt.Println(&n, n)
		// #2 here `n` is reassigned or declared? former
		n, m := 10, 100
		fmt.Println(&n, n)

		// no new variables on left side of :=
		// n, _ := 100, 1000
		// fmt.Println(&n, n)

		// #1 and #2 are the same? YES
		_ = m
	}

	{
		// #1
		var n int

		if a := os.Args; len(a) == 1 {
			fmt.Println("a[0]:", a[0])
		} else if n, e := strconv.Atoi(a[1]); e == nil { // #2 here `n` is reassigned or declared? latter
			fmt.Println(&n, n)
		} else {
			fmt.Println("a[1]:", a[1])
			fmt.Println("e:", e)
		}

		// #1 and #2 are the same? NO
		fmt.Println(&n, n)
	}

}
