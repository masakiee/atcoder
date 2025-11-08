package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	var ans int = ((B-1) + (A-1) - 1) / (A-1)
	fmt.Println(ans)
}
