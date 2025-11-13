package main

import (
	"fmt"
)

func main() {
	var A, B, C, D int
	fmt.Scanf("%d %d %d %d ", &A, &B, &C, &D)

	tAtk := (C+B-1)/B
	aAtk := (A+D-1)/D
	if tAtk <= aAtk {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
