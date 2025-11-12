package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	sum := sumInts(A...)
	count := 0
	for _, a := range A {
		if a >= (sum+4*M-1)/(4*M) {
			count++
		}
	}
	
	if count >= M {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
