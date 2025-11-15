package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	A := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sort.Ints(A)

	iter := slices.Backward(A)
	for _, i := range iter {
		fmt.Print(i)
	}
}
