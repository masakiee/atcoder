package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	P := make([]int, N)
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &P[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &Q[i])
	}

	per := []int{}
	for i:=1; i<=N; i++ {
		per = append(per, i)
	}

	var a, b int
	for i:=0; i<factorial(N); i++ {
		pEqual, qEqual := true, true
		for j:=0; j<N; j++ {
			if P[j] != per[j] {
				pEqual = false
			}
			if Q[j] != per[j] {
				qEqual = false
			}
		}

		if pEqual {
			a = i
		}
		if qEqual {
			b = i
		}

		per = nextPermutation(per)
	}

	fmt.Println(absInt(a - b))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func nextPermutation(x []int) []int {
	var n = len(x)
	var left, right int

	for i := n-2; i >= 0; i-- {
		if x[i] < x[i+1] {
			left = i
			break
		}
	}
	for i := n-1; i >= 0; i-- {
		if x[left] < x[i] {
			right = i
			break
		}
	}

	x[left], x[right] = x[right], x[left]

	for i := 0; i < (n-left)/2; i++ {
		x[left+1+i], x[n-1-i] = x[n-1-i], x[left+1+i]
	}

	return x
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
