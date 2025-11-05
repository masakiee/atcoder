package main

import (
	"fmt"
	"math"
)

type Gem struct {
	Value int
	Cost int
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	V := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &V[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &C[i])
	}

	var ans int = math.MinInt
	for mask := 0; mask < (1 << N); mask++ {
		var cp int
		for i := 0; i < N; i++ {
			if mask & (1<<i) != 0 {
				cp += V[i] - C[i]
			}
		}
		ans = max(ans, cp)
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
