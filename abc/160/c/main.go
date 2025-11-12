package main

import (
	"fmt"
	"sort"
)

func main() {
	var K, N int
	fmt.Scan(&K)
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	D := make([]int, N)
	D[0] = A[0] + (K - A[N-1])
	for i := 1; i < N; i++ {
		D[i] = A[i] - A[i-1]
	}

	sort.Ints(D)
	
	var ans int = sumInts(D[0:N-1]...)
	fmt.Println(ans)
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
