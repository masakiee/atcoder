package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)
	A := make([]int, M)
	re := make(map[int]bool)
	for i := range A {
		fmt.Scan(&A[i])
		re[A[i]] = true
	}
	
	ans := []int{}
	L := 1
	for L <= N {
		R := L
		for re[R] {
			R++
		}

		for i:=R; i>=L; i-- {
			ans = append(ans, i)
		}

		L = R + 1
	}

	for _, v := range ans {
		fmt.Print(v, " ")
	}
}
