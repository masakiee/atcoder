package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)

	var ans int = max(C - (A - B), 0)
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
