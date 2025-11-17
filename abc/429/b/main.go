package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)
	A := make([]int, N)
	aMap := make(map[int]struct{})
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		aMap[A[i]] = struct{}{}
	}
	sub := sumInts(A...) - M
	_, ok := aMap[sub]
	if ok {
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
