package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	ans := minInts(b/a, c)
	fmt.Println(ans)
}

func minInts(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
