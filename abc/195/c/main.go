package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	ans := 0
	MAX_N := 1_000_000_000_000_000
	for i := 1_000; i <= MAX_N; i *= 1_000 {
		ans += max(0, N - (i-1))
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
