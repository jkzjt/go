package main

func main() {
	// ------------------------------------
	// constants are immutable
	// ------------------------------------

	// const n = 10
	// build fail: cannot assign to n
	// n = 100

	// build fail: a (variable of type int) is not constant
	// a := 1
	// const m = a

	const n = len("Hi") // allowed

	// build fail
	// s := "元宇宙"
	// const n = len(s)

	const i = 100
	var f float64 = i

	_, _, _ = n, i, f
}
