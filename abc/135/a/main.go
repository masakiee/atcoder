package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	k2 := A + B
	if k2 % 2 == 1 {
		fmt.Println("IMPOSSIBLE")
		return
	}

	var ans int = k2 / 2
	fmt.Println(ans)
}
