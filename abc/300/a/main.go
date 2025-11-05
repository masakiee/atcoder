package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scanf("%d %d %d", &N, &A, &B)
	sum := A + B
	C := make([]int, N)
	var ans int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &C[i])
		if C[i] == sum {
			ans = i+1
		}
	}
	
	fmt.Println(ans)
}
