package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	counts := make(map[int]int)
	for i := range A {
		fmt.Scanf("%d", &A[i])
		m := A[i]%200
		counts[m]++
	}
	
	ans := 0
	for _, cnt := range counts {
		if cnt > 1 {
			ans += nPr(cnt, 2) / 2
		}
	}
	fmt.Println(ans)
}

func nPr(n, r int) int {
	v := 1
	for i:=n; i>n-r; i-- {
		v *= i
	}
	return v
}
