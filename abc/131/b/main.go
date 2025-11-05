package main

import (
	"fmt"
	"math"
)

func main() {
	var N, L int
	fmt.Scanf("%d %d", &N, &L)


	A := make([]int, N)
	minAji := math.MaxInt
	for i := 1; i <= N; i++ {
		A[i-1] = L + i - 1
		if absInt(minAji) > absInt(A[i-1]) {
			minAji = A[i-1]
		}
	}

	ajiSum := sumInts(A...)
	ans := ajiSum - minAji
	
	fmt.Println(ans)
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
