package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	fmt.Println(solve(S))
}

func solve(S string) string {
	if strings.Contains(S, "ab") || strings.Contains(S, "ba") {
		return "Yes"
	} else {
		return "No"
	}
}
