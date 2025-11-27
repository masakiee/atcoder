package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make(map[int]int)
	for i:=0; i<N; i++ {
		var a int
		fmt.Scan(&a)
		A[a]++
	}
	B := make([]int, N+1)
	for i:=1; i<=N; i++ {
		fmt.Scan(&B[i])
	}
	C := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&C[i])
	}

	BC := make(map[int]int)
	for _, Cj := range C {
		BC[B[Cj]]++
	}

	var ans int
	for Ai, cntA := range A {
		if cntB, ok := BC[Ai]; ok {
			ans += cntA * cntB
		}
	}
	fmt.Println(ans)
}
