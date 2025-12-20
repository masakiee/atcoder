package main

import (
	"fmt"
)

func main() {
	var h, w, n int
	fmt.Scanf("%d %d %d", &h, &w, &n)

	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
		for j := 0; j < w; j++ {
			var v int
			fmt.Scan(&v)
			m[i][j] = v
		}
	}

	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		for j := 0; j < h; j++ {
			for _, value := range m[j] {
				if value == v {
					cnt[j]++
					break
				}
			}
		}
	}

	ans := 0
	for _, v := range cnt {
		if v > ans {
			ans = v
		}
	}

	fmt.Println(ans)
}
