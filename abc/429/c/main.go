package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		B[A[i]-1]++
	}
	
	var ans int
	for _, b := range B {
		ans += b*(b-1)/2 * (N-b)
	}

	fmt.Println(ans)
}
