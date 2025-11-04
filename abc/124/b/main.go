package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &H[i])
	}

	memo := make(map[int]struct{})
	max := H[0]
	for i, h := range H {
		if h >= max {
			max = h
			memo[i] = struct{}{}
		}
	}

	fmt.Println(len(memo))
}
