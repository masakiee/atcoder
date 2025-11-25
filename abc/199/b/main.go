package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	B := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	for i := range B {
		fmt.Scanf("%d", &B[i])
	}

	seen := make(map[int]int)
	for i := 0; i < N; i++ {
		for j := A[i]; j <= B[i]; j++ {
			seen[j]++
		}
	}

	ans := 0
	for _, size := range seen {
		if size == N {
			ans++
		}
	}
	fmt.Println(ans)
}
