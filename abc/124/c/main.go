package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	headBlack := flipCount(S, '0')
	headWhite := flipCount(S, '1')

	fmt.Println(min(headBlack, headWhite))
}

func flipCount(s string, headColor rune) int {
	cnt := 0
	for i, c := range s {
		if i % 2 == 0 {
			if headColor != c {
				cnt++
			}
		} else {
			if headColor == c {
				cnt++
			}
		}
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
