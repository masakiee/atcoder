package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	
	fmt.Println(solve(S))
}

func solve(S string) string {
	for i, s := range strings.Split(S, "") {
		if (i+1) % 2 == 1 {
			if isUpper(s) {
				return "No"
			}
		} else {
			if isLower(s) {
				return "No"
			}
		}
	}
	return "Yes"
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}
func isLower(s string) bool {
	return strings.ToLower(s) == s
}
