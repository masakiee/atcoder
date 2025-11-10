package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	
	var ans int
	fmt.Println(ans)
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
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

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}

func lowerBound(a []int, v int) int {
	low, high := -1, len(a)
	for high - low > 1 {
		mid := (high - low) / 2
		if a[mid] >= v {
			high = mid
		} else {
			low = mid
		}
	}

	return high
}

func upperBound(a []int, v int) int {
	low, high := -1, len(a)
	for high - low > 1 {
		mid := (high + low) / 2
		if a[mid] > v {
			high = mid
		} else {
			low = mid
		}
	}
	return high
}

func ceil(a, b int) int {
	return (a + b - 1) / b
}

func filter[T any](arr []T, f func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func numPer(n int) int {
	ans := 1
	for i:=n; i>0; i-- {
		ans *= n
	}
	return ans
}
