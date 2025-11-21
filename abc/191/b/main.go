package main

import (
	"fmt"
)

func main() {
	var N, X int
	fmt.Scan(&N)
	fmt.Scan(&X)
	A := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	for _, v := range A {
		if v != X {
			fmt.Print(v, " ")
		}
	}
}
