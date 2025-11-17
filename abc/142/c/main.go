package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	type Record struct {
		i, order int
	}
	A := make([]Record, N)
	for i := 0; i < N; i++ {
		var a int
		fmt.Scanf("%d", &a)
		A[i] = Record{i+1, a}
	}

	sort.Slice(A, func (i, j int) bool {
		return A[i].order < A[j].order
	})

	for i, r := range A {
		if i == 0 {
			fmt.Print(r.i)
		} else {
			fmt.Printf(" %d", r.i)
		}
	}
}
