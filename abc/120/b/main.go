package main

import (
	"fmt"
)

func main() {
	var A, B, K int
	fmt.Scanf("%d %d %d", &A, &B, &K)

	div := []int{}
	for i := 100; i >= 1; i-- {
		if A % i == 0 && B % i == 0 {
			div = append(div, i)
			if len(div) == K {
				fmt.Println(i)
				return
			}
		}
	}	
}
