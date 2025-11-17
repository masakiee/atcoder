package main

import (
	"fmt"
)

func main() {
	var N, K, M int
	fmt.Scanf("%d %d %d", &N, &K, &M)
	A := make([]int, N)
	for i := 0; i < N-1; i++ {
		fmt.Scanf("%d", &A[i])
	}
	sum := sumInts(A...)
	diff := M*N - sum
	if diff > K {
		fmt.Println(-1)
	} else {
		fmt.Println(maxInts(0, diff))
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
