package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scanf("%d %d %d", &A, &B, &C)
	fmt.Println(solve(A, B, C))
}

func solve(A, B, C int) string {
	takahashi := "Takahashi"
	aoki := "Aoki"
	if C == 0 {
		if A > B {
			return takahashi
		}  else {
			return aoki
		}
	} else {
		if B > A {
			return aoki
		} else {
			return takahashi
		}
	}
}
