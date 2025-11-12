package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	
	var ans int = minInts(N%K, K-N%K)
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
