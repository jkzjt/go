package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	tip   = "Need a magnitude"
	haze  = "Couldn't get it"
	msg   = "%g is %s"
	wrong = "%g is unreasonable"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}
	s := strings.TrimSpace(os.Args[1])
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(haze)
		return
	}
	switch {
	case f >= 10.:
		fmt.Printf(msg, f, "massive")
	case f >= 8.:
		fmt.Printf(msg, f, "great")
	case f >= 7.:
		fmt.Printf(msg, f, "major")
	case f >= 6.:
		fmt.Printf(msg, f, "strong")
	case f >= 5.:
		fmt.Printf(msg, f, "moderate")
	case f >= 4.:
		fmt.Printf(msg, f, "light")
	case f >= 3.:
		fmt.Printf(msg, f, "minor")
	case f >= 2.:
		fmt.Printf(msg, f, "very minor")
	case f > 0.:
		fmt.Printf(msg, f, "micro")
	default:
		fmt.Printf(wrong, f)
	}
}
