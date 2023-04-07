package main

import "fmt"

func main() {
	var (
		e          byte    = 'T'       // an english letter
		r          rune    = '谭'       // a non-english letter
		y          uint16  = 2050      // a year in 4-digit like 2050
		m          uint8   = 9         // a month in 2-digit without leading zeros
		lightSpeed uint32  = 299792458 // the speed of light in `m/s`
		angle      float32 = 34.8      // angle of circle
		pi         float64 = 3.1415926 // PI
	)
	fmt.Println("                       an english letter:", e)
	fmt.Println("                    a non-english letter:", r)
	fmt.Println("             a year in 4-digit like 2050:", y)
	fmt.Println("a month in 2-digit without leading zeros:", m)
	fmt.Println("             the speed of light in `m/s`:", lightSpeed)
	fmt.Println("                         angle of circle:", angle)
	fmt.Println("                                      PI:", pi)
}
