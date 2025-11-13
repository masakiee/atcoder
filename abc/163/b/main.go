package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)
	A := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &A[i])
	}
	sum := sumInts(A...)
	if sum > N {
		fmt.Println(-1)
	} else {
		fmt.Println(N - sum)
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
