package main

// -----------------------
// VARIABLE NAMING RULES
// -----------------------
func main() {
	// allowed
	var speed float64
	var Speed float64

	// underscore `_` allowed but not recommended
	var _speed float64

	// Unicode letters allowed
	var 速度 float64

	_, _, _, _ = speed, Speed, _speed, 速度
}
