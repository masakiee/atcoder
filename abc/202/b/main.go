package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	m := map[string]string {
		"0": "0",
		"1": "1",
		"6": "9",
		"8": "8",
		"9": "6",
	}
	var ans string
	for _, c := range strings.Split(S, "") {
		ans = m[c] + ans
	}
	
	fmt.Println(ans)
}
