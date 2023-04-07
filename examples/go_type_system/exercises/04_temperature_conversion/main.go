package main

import "fmt"

type (
	Temperature float64
	Celsius     Temperature
	Fahrenheit  Temperature
)

func main() {
	var (
		factor                 = 2.
		celDegree  Temperature = 10.5
		fahDegree  Temperature = 5.5
		celsius    Celsius     = 69.8
		fahrenheit Fahrenheit  = 18.6
	)

	celsius *= Celsius(float64(celDegree) * factor)
	fahrenheit *= Fahrenheit(float64(fahDegree) * factor)
	fmt.Println(celsius, fahrenheit)
}
