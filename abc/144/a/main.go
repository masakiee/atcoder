package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	
	if A >= 10 || B >= 10 {
		fmt.Println(-1)
	} else {
		fmt.Println(A * B)
	}
}
