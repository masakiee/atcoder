package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	ans := A*12 + B
	fmt.Println(ans)
}
