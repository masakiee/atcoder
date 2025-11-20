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

	var ans int
	sum := sumInts(A...)
	for _, a := range A {
		ans += a * a * (N-1)
		ans -= a * (sum - a)
	}
	fmt.Println(ans)
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
