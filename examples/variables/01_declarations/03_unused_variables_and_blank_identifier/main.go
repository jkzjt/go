package main

// no warnning for package-level variables
var variableInPackageLevel int

func main() {
	// unused -- it will yield an error in compile-time
	var tmp int // tmp declared but not used

	// when used the error will be gone
	// print(tmp)
	
	// we can assign the variable to the blank identifier
	// for passing compile-time and everything is ok
	_ = tmp
}
