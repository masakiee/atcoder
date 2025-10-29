package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		B[i] = A[i] - i
	}
	sort.Slice(B, func (i, j int) bool {
		return B[i] < B[j]
	})

	var median int
	if N % 2 == 1 {
		median = B[N/2]
	} else {
		median = (B[N/2-1] + B[N/2]) / 2
	}

	var ans int
	for _, bi := range B {
		if bi - median >= 0 {
			ans += bi - median
		} else {
			ans -= bi - median
		}
	}
	
	out(ans)
}

func out(x ...interface{}) {
	fmt.Println(x...)
}
