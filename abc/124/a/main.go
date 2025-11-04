package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	var ans int
	for i:=0; i<2; i++ {
		if A > B {
			ans += A
			A--
		} else {
			ans += B
			B--
		}
	}

	fmt.Println(ans)
}
