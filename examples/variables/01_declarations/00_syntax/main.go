package main

import (
	"fmt"
)

func main() {
	// var -- keyword for declare variables in GO

	//-----------------------------------
	// syntax:
	//		var variable_name data_type
	//-----------------------------------

	// examples
	// int
	var nFiles int
	var nCPU int
	var level int

	fmt.Println(
		nFiles,
		nCPU,
		level,
	)

	// float64
	var temperature float64
	var humidity float64
	var ratio float64

	fmt.Println(
		temperature,
		humidity,
		ratio,
	)

	// bool constants
	var allowed bool
	var on bool
	var closed bool

	fmt.Println(
		allowed,
		on,
		closed,
	)

	// string
	var username string
	var password string
	var salt string

	fmt.Println(
		username,
		password,
		salt,
	)
}
