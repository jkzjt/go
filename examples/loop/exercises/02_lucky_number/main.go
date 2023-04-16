package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	tip      = "Pick a number"
	wrongArg = "Wrong argument: %q\n"
	lose     = "Defeat"
	win      = "Victory"
	maxTurns = 5
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(tip)
		return
	}

	s := strings.TrimSpace(os.Args[1])
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf(wrongArg, s)
		return
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	fmt.Print("[")
	var (
		num int
		msg = lose
	)
	for i := 1; i <= maxTurns; i++ {
		num = r.Intn(n + 1)
		if num == n {
			fmt.Print(num)
			msg = win
			break
		} else if i != maxTurns {
			fmt.Print(num, " ")
		} else {
			fmt.Print(num)
		}
	}
	fmt.Println("]")
	fmt.Println(msg)
}
