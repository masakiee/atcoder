package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	diffs := make([]int, N-1)
	for i:=0; i<N-1; i++ {
		diffs[i] = A[i+1]-A[i]
	}

	var newArr []int = []int{A[0]}
	for _, diff := range diffs {
		start := newArr[len(newArr)-1]
		delta := diff/absInt(diff)
		for i:=1; i<=absInt(diff); i++ {
			newArr = append(newArr, start+delta*i)
		}
	}
	ans := fmt.Sprintf("%v", newArr)
	ans = ans[1:len(ans)-1]
	fmt.Println(ans)
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
