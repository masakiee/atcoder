package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	A := make([]int, M)
	B := make([]int, M)
	for i:=0; i<M; i++ {
		fmt.Scanf("%d %d", &A[i], &B[i])
	}
	var K int
	fmt.Scan(&K)
	C := make([]int, K)
	D := make([]int, K)
	for i:=0; i<K; i++ {
		fmt.Scanf("%d %d", &C[i], &D[i])
	}

	ans := 0
	for mask:=0; mask<1<<K; mask++ {
		ballDish := map[int]bool{}
		for i:=0; i<K; i++ {
			if mask & (1<<i) == 0 {
				// c を選ぶ
				ballDish[C[i]] = true
			} else {
				// d を選ぶ
				ballDish[D[i]] = true
			}
		}

		// 条件は高々M個なのでチェック
		cnt := 0
		for i:=0; i<M; i++ {
			if ballDish[A[i]] && ballDish[B[i]] {
				cnt++
			}
		}
		ans = maxInts(ans, cnt)
	}

	fmt.Println(ans)
}

func maxInts(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
