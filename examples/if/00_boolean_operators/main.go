package main

import . "fmt"

func main() {
	// ----------------------------------------------
	// COMPARISON OPERATORS
	// >, <, >=, <=
	// ABOVE ONLY WORK WITH NUMERIC VALUES
	// ==, !=
	// ----------------------------------------------
	{
		const (
			pass = iota*10 + 60
			fine
			excellent
			best
			perfect
		)
		mark, valid := 79, true
		Println("valid?", valid)
		Println("perfect?", mark == perfect)
		Println("best?", mark >= best)
		Println("excellent?", mark >= excellent)
		Println("fine?", mark >= fine)
		Println("pass?", mark >= pass)
		Println("fail?", mark < pass)

		const (
			low = iota*30 + 60
			mid
			high
		)
		speed := 100
		death := speed > high
		fast := speed > mid
		normal := speed == mid
		slow := speed <= low
		Println("death?", death)
		Println("fast?", fast)
		Println("normal?", normal)
		Println("slow?", slow)
	}

	// --------------------------------------
	// LOGICAL OPERATORS
	// AND(&&) OR(||) NOT(!)
	// ONLY WORK WITH BOOLEAN EXPRESSION
	// --------------------------------------
	{
		// &&
		Println("true && true =", true && true)
		Println("true && false =", true && false)
		Println("false && true =", false && true)
		Println("false && false =", false && false)
		// ||
		Println("true || true =", true || true)
		Println("true || false =", true || false)
		Println("false || true =", false || true)
		Println("false || false =", false || false)
		// !
		Println("!true =", !true)
		Println("!false =", !false)
		Println("!(false || true) =", !(false || true))
	}

	{
		const (
			pass = iota*10 + 60
			fine
			excellent
			best
			perfect
		)
		mark := 79
		Println("perfect?", mark == perfect)
		Println("best?", mark >= best && mark < perfect)
		Println("excellent?", mark >= excellent && mark < best)
		Println("fine?", mark >= fine && mark < excellent)
		Println("pass?", mark >= pass && mark < fine)
		Println("fail?", mark < pass)

		const (
			low = iota*30 + 60
			mid
			high
		)
		speed := 100
		death := speed > high
		fast := speed > mid && speed <= high
		normal := speed <= mid && speed > low
		slow := speed <= low
		Println("death?", death)
		Println("fast?", fast)
		Println("normal?", normal)
		Println("slow?", slow)
	}
}
