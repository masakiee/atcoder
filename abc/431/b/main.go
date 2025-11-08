package main

import (
	"fmt"
)

func main() {
	var X, N, Q int
	fmt.Scanf("%d", &X)
	fmt.Scanf("%d", &N)
	W := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &W[i])
	}
	fmt.Scanf("%d", &Q)

	equipped := map[int]bool{}
	P := make([]int, Q)
	S := []int{}
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d", &P[i])
		equipped[P[i]] = !equipped[P[i]]
		wSum := X
		for i, equipped := range equipped {
			if equipped {
				wSum += W[i-1]
			}
		}
		S = append(S, wSum)
	}

	for _, s := range S {
		fmt.Println(s)
	}
}

