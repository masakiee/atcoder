package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	AS := make([][]int, M)
	for i := range AS {
		var c int
		fmt.Scan(&c)
		AS[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&AS[i][j])
		}
	}
	
	ans := 0
	for mask := 1; mask < 1<<M; mask++ {
		num := map[int]bool{}
		for i := 0; i < M; i++ {
			if mask & (1 << i) != 0 {
				for _, v := range AS[i] {
					num[v] = true
				}
			}
		}
		if len(num) == N {
			ans++
		}
	}
	fmt.Println(ans)
}
