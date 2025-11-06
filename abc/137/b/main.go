package main

import (
	"fmt"
)


const MIN_X int = -1_000_000
const MAX_X int =  1_000_000

func main() {
	var K, X int
	fmt.Scanf("%d", &K)
	fmt.Scanf("%d", &X)
	
	startX := maxInts(MIN_X, X - K + 1)
	endX   := minInts(MAX_X, X + K - 1)

	ans := []int{}
	for x := startX; x <= endX; x++ {
		ans = append(ans, x)
	}

	for i, v := range ans {
		if i == 0 {
			fmt.Print(v)
		} else {
			fmt.Printf(" %d", v)
		}
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

func minInts(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
	}
	return min
}
