package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &K)

	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &H[i])
	}

	filtered := filter(H, func (h int) bool {
		return h >= K
	})
	
	var ans int = len(filtered)
	fmt.Println(ans)
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
