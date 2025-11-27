package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)

	for i:=1; i<=K; i++ {
		if N % 200 == 0 {
			N /= 200
		} else {
			ns := strconv.Itoa(N) + "200"
			N, _ = strconv.Atoi(ns)
		}
	}

	var ans int = N
	fmt.Println(ans)
}
