package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}

	ans := 0
	for l:=0; l<N; l++ {
		x := A[l]
		for r:=l; r<N; r++ {
			x = minInts(x, A[r])
			ans = maxInts(ans, x*(r-l+1))
		}
	}
	
	fmt.Println(ans)
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
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
