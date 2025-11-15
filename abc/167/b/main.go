package main

import (
	"fmt"
)

func main() {
	var A, B, C, K int
	fmt.Scanf("%d %d %d %d", &A, &B, &C, &K)

	var ans int
	if A >= K {
		ans = K
	} else if A + B >= K {
		ans = A
	} else {
		ans = A + (-1) * (K - (A + B))
	}
	
	fmt.Println(ans)
}
