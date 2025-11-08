package main

import (
	"fmt"
)

func main() {
	var N, M, K int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)
	fmt.Scanf("%d", &K)

	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &H[i])
	}
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &B[i])
	}

	aParts := H
	bParts := B
	isBaseHead := true
	if len(B) > len(H) {
		aParts = B
		bParts = H
		isBaseHead = false
	}

	// 累積和
	cnt := map[int]int{}
	for _, w := range bParts {
		cnt[w]++
	}

	num := 0
	for _, wa := range aParts {
		var ok bool
		if isBaseHead {
			ok = cumsum(cnt, 200000) - cumsum(cnt, wa-1) > 0
		} else {
			ok = cumsum(cnt, wa) > 0
		}

		if ok {
			num++
		}
	}
	ok := num >= K
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func cumsum(m map[int]int, v int) int {
	s := 0
	for value, cnt := range m {
		if value <= v {
			s += cnt
		}
	}
	return s
}
