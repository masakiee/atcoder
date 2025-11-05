package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	
	var ans int = minInts(B, N*A)
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
