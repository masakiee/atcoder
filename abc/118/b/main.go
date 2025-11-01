package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	ret := make(map[int]int)
	for i := 0; i < N; i++ {
		var K int
		fmt.Scanf("%d", &K)

		for j := 0; j < K; j++ {
			var a int
			fmt.Scanf("%d", &a)
			ret[a]++
		}
	}

	cnt := 0
	for _, num := range ret {
		if num == N {
			cnt++
		}
	}
	
	fmt.Println(cnt)
}
