package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N-1)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &B[i])
	}
	for i := 0; i < N-1; i++ {
		fmt.Scanf("%d", &C[i])
	}
	
	var ans int
	for i, a := range A {
		ans += B[a-1]
		if i >= 1 && A[i]-A[i-1] == 1 {
			ans += C[a-1-1]
		}
	}
	fmt.Println(ans)
}
