package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	honests := map[int][]int{}
	unkinds := map[int][]int{}
	for i := 0; i < N; i++ {
		var A int
		fmt.Scanf("%d", &A)

		honests[i] = []int{}
		unkinds[i] = []int{}
		for j:=0; j<A; j++ {
			var x, y int
			fmt.Scanf("%d %d", &x, &y)
			if y == 1 {
				honests[i] = append(honests[i], x-1)
			} else {
				unkinds[i] = append(unkinds[i], x-1)
			}
		}
	}

	var ans int
	for mask := 0; mask < 1<<N; mask++ {
		ok := true
		total := 0
		for i:=0; i<N; i++ {
			if mask & (1 << i) == 0 {
				continue
			}
			total++
			for _, honest := range honests[i] {
				if mask & (1 << honest) == 0 {
					ok = false
				}
			}
			for _, unkind := range unkinds[i] {
				if mask & (1 << unkind) != 0 {
					ok = false
				}
			}
		}
		if ok {
			ans = maxInts(ans, total)
		}
	}
	
	fmt.Println(ans)
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
