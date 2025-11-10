package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sort.Ints(A)
	alice := 0
	bob := 0
	for i:=0; i<N; i++ {
		v := A[N-1-i]
		if i % 2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}
	
	var ans int = alice - bob
	fmt.Println(ans)
}
