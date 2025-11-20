package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)
	
	fmt.Println(solve(A, B))
}

func solve(A, B int) int {
	C := A + B
	switch true {
	case C >= 15 && B >= 8:
		return 1
	case C >= 10 && B >= 3:
		return 2
	case C >= 3:
		return 3
	default:
		return 4
	}
}
