package main

import (
	"fmt"
)

func main() {
	var S, A, B, X int
	fmt.Scanf("%d %d %d %d", &S, &A, &B, &X)


	var ans int
	set := X / (A + B)
	ans += S * A * set
	ans += S * minInts(A, X % (A + B))
	
	fmt.Println(ans)
}

func minInts(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
	}
	return min
}
