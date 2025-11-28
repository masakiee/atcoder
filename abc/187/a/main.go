package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	
	ans := max(S(A), S(B))
	fmt.Println(ans)
}

func S(n int) int {
	ans := 0
	d := 1
	for n / d > 0 {
		d *= 10
		ans += (n%d) / (d/10)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
