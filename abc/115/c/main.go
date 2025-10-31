package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &H[i])
	}

	sort.Ints(H)

	diffs := []int{}
	for i := 0; i <= N-K; i++ {
		h := H[i:i+K]
		hmin := h[0]
		hmax := h[len(h)-1]
		diffs = append(diffs, hmax-hmin)
	}
	fmt.Println(min(diffs))
}


func min(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if m > v {
			m = v
		}
	}

	return m
}
