package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		B[i] = A[i] / 2
	}

	lc := 1
	for i := 0; i < N; i++ {
		lc = lcm(lc, B[i])
	}
	for i := 0; i < N; i++ {
		if (lc / B[i]) % 2 == 0 {
			fmt.Println(0)
			return
		}
	}

	ans := (M/lc + 2 - 1) / 2
	fmt.Println(ans)
}

func gcm(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcm(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / gcm(a, b)
}
