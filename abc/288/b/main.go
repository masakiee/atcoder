package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	S := make([]string, N)
	for i := range S {
		fmt.Scan(&S[i])
	}

	S = S[0:K]
	sort.Slice(S, func (i, j int) bool {
		return S[i] < S[j]
	})
	
	for _, v := range S {
		fmt.Println(v)
	}
}
