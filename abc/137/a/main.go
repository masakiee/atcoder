package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	add := A + B
	sub := A - B
	mul := A * B
	var ans int = maxInts(add, sub, mul)
	fmt.Println(ans)
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
