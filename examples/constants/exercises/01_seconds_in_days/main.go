package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Duration int64

const (
	secondsInAMinute Duration = 60
	minutesInAnHour
	hoursInADay Duration = 24
	usage                = `
This program tells certain seconds in days you gave

USAGE: 
  ./main.exe days`
	illegalArg = "Illegal argument: %q\n"
	tpl        = "There are %d seconds in %d days\n"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}
	day := strings.TrimSpace(os.Args[1])
	if v, err := strconv.ParseInt(day, 10, 64); err == nil {
		d := Duration(v)
		totalSeconds := d * hoursInADay * minutesInAnHour * secondsInAMinute
		fmt.Printf(tpl, totalSeconds, d)
		return
	}
	fmt.Printf(illegalArg, day)
}
