package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	pass = iota*10. + 60.
	fine
	excellent
	best
	perfect
)

const (
	FAIL      = "FAIL"
	PASS      = "PASS"
	FINE      = "FINE"
	EXCELLENT = "EXCELLENT"
	BEST      = "BEST"
	PERFECT   = "PERFECT"
)

const (
	usage = `
-------------------------------
 This program for score remark
-------------------------------
USAGE: 
  ./main.exe score0 score1 ...`
	tpl          = "%q -- %q\n"
	illegalArg   = "Illegal argument: %q\n"
	invalidScore = "Invalid score: %q\n"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}
	var remark string
	for _, v := range os.Args[1:] {
		v = strings.TrimSpace(v)
		if d, err := strconv.ParseFloat(v, 64); err == nil {
			if d < 0 || d > 100. {
				fmt.Printf(invalidScore, v)
				continue
			} else if d == perfect {
				remark = PERFECT
			} else if d >= best {
				remark = BEST
			} else if d >= excellent {
				remark = EXCELLENT
			} else if d >= fine {
				remark = FINE
			} else if d >= pass {
				remark = PASS
			} else {
				remark = FAIL
			}
			fmt.Printf(tpl, v, remark)
		} else {
			fmt.Printf(illegalArg, v)
		}
	}
}
