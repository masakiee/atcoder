package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	p := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &p[i])
	}

	sort.Ints(p)

	var sum int
	for i, v := range p {
		if i == len(p)-1 {
			sum += v/2
		} else {
			sum += v
		}
	}
	
	fmt.Println(sum)
}
