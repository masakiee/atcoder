package main

import (
	"fmt"
)

func main() {
	A := make([]int, 3)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	
	fmt.Println(solve(A))
}

func solve(A []int) string {
	switch true {
	case A[0]-A[1] == A[1]-A[2]:
		return "Yes"
	case A[1]-A[2] == A[2]-A[0]:
		return "Yes"
	case A[2]-A[0] == A[0]-A[1]:
		return "Yes"
	default:
		return "No"
	}
}
