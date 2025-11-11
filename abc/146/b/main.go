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

	var ans string
	for _, s := range strings.Split(S, "") {
		ans += rot(s, N)
	}
	
	fmt.Println(ans)
}

var abc string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
func rot(s string, n int) string {
	idx := strings.Index(abc, s)
	idx = (idx + n) % len(abc)
	return string(abc[idx])
}
