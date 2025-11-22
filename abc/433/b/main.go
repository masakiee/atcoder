package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N+1)
	for i:=1; i<=N; i++ {
		fmt.Scanf("%d", &A[i])
		idx := -1
		for j:=i-1; j>=1; j-- {
			if A[j] > A[i] {
				idx = j
				break
			}
		}
		fmt.Println(idx)
	}
}
