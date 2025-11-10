package main

import (
	"fmt"
)

func main() {
	var A, B, C, X int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)
	
	ans := 0
	for i:=0; i<=A; i++ {
		for j:=0; j<=B; j++ {
			for k:=0; k<=C; k++ {
				sum := 500*i + 100*j + 50*k
				if sum == X {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
