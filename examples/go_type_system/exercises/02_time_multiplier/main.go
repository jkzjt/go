package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	h, err := time.ParseDuration("4h30m")
	if err != nil {
		return
	}
	m, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		return
	}
	h *= time.Duration(m)
	fmt.Println(h)
}
