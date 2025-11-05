package main

import (
	"fmt"
)

func main() {
	var A, P int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &P)

	pSum := A * 3 + P
	ans := pSum / 2
	fmt.Println(ans)
}
