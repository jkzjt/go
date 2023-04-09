package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Feet  float64
	Meter float64
)

const (
	factor = 0.3048
	usage  = `
This program is for converting feet to meters

USAGE: 
	./main.exe feet`
	illegalArg = "Illegal argument: %q\n"
	tpl        = "%g Feet is %.2f Meters\n"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}
	if v, err := strconv.ParseFloat(strings.TrimSpace(os.Args[1]), 64); err == nil {
		feet := Feet(v)
		meters := Meter(feet) * factor
		fmt.Printf(tpl, feet, meters)
		return
	}
	fmt.Printf(illegalArg, strings.TrimSpace(os.Args[1]))
}
