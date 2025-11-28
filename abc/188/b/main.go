package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	B := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	var ans int
	for i := range B {
		fmt.Scanf("%d", &B[i])
		ans += A[i]*B[i]
	}

	if ans == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
