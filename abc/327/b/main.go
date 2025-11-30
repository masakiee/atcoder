package main

import (
	"fmt"
)

func main() {
	var B int
	fmt.Scan(&B)

	for A:=1; A<=15; A++ {
		if pow(A, A) == B {
			fmt.Println(A)
			return
		}
	}

	fmt.Println(-1)
}

func pow(a, b int) int {
	s := 1
	for i:=0; i<b; i++ {
		s *= a
	}
	return s
}
