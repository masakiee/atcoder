package main

import (
	"fmt"
)

func main() {
	var N, X int
	fmt.Scan(&N)
	fmt.Scan(&X)
	V := make([]int, N)
	P := make([]int, N)
	for i := range V {
		fmt.Scanf("%d %d", &V[i], &P[i])
	}

	sum := 0
	for i := range V {
		sum += V[i] * P[i]
		if sum > X * 100 {
			fmt.Println(i+1)
			return
		}
	}
	
	fmt.Println(-1)
}
