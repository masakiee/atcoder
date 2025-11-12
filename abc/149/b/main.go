package main

import (
	"fmt"
)

func main() {
	var A, B, K int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&K)

	if A - K >= 0 {
		fmt.Printf("%d %d", A-K, B)
	} else {
		fmt.Printf("%d %d", 0, maxInts(0, B-(K-A)))
	}
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
