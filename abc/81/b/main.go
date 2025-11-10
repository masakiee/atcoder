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

	ans := 0
	for valid(A) {
		for i := range A {
			A[i] /= 2
		}
		ans++
	}
	
	fmt.Println(ans)
}
func valid(arr []int) bool {
	for i := range arr {
		if arr[i] % 2 == 1 {
			return false
		}
	}
	return true
}
