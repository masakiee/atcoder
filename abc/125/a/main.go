package main

import (
	"fmt"
)

func main() {
	var A, B, T int
	fmt.Scanf("%d %d %d", &A, &B, &T)

	var ans int
	for t := A; t <= T; t += A {
		ans += B
	}
	
	fmt.Println(ans)
}
