package main

import (
	"fmt"
)


func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	s := (A-1) * (B-1)
	fmt.Println(s)
}
