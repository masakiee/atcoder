package main

import (
	"fmt"
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

	cp := make([]int, N)
	for i := range V {
		cp[i] = V[i] - C[i]
	}

	cp = filter(cp, func (v int) bool {
		return v > 0
	})
	
	fmt.Println(sum(cp))
}

func sum(arr []int) int {
	var ret int
	for _, v := range arr {
		ret += v
	}
	return ret
}

func filter[T any](arr []T, f func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
