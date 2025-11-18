package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := map[string]string{
		"0": "1",
		"1": "0",
	}
	var ans string
	for _, c := range strings.Split(s, "") {
		ans += t[c]
	}
	
	fmt.Println(ans)
}
