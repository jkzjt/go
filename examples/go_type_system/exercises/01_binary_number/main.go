package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// decimalToBinary converts an integer typed int64 to s string with binary form
func decimalToBinary(d int64) (b string) {
	// a comparable number for eliminating leading zeros
	var n float64 = 0

	for i := 0.; i < 64.; i++ {
		if d>>int(i)&1 == 1 {
			b = "1" + b
			n += math.Pow(2, i)
			// there is no need to go on if satisfied
			// or `b` may has leading zeros
			if int64(n) == d {
				break
			}
		} else {
			b = "0" + b
		}
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	d, err := strconv.ParseInt(strings.TrimSpace(os.Args[1]), 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("%q = %q\n", os.Args[1], decimalToBinary(d))
}
