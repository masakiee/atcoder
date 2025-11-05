package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scanf("%d", &N)
	fmt.Scan(&S)
	tWin := 0
	aWin := 0
	minWin := (N+2-1)/2
	for _, ret := range strings.Split(S, "") {
		if ret == "T" {
			tWin++
		} else {
			aWin++
		}

		if tWin >= minWin {
			fmt.Println("T")
			return
		}
		if aWin >= minWin {
			fmt.Println("A")
			return
		}
	}
}
