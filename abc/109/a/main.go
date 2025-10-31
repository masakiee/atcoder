package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	if (A * B) % 2 == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
