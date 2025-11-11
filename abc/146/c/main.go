package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B, X int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&X)

	var ok int = 0
	var ng int = 1_000_000_000 + 1
	for ok + 1 != ng {
		md := (ok + ng) / 2
		if price(A, B, md) > X {
			ng = md
		} else {
			ok = md
		}
	}
	
	var ans int = ok
	fmt.Println(ans)
}

func price(A, B, N int) int {
	return A * N + B * d(N)
}

func d(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}
