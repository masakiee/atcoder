package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	D := make([]int, N)
	set := map[int]struct{}{}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &D[i])
		set[D[i]] = struct{}{}
	}
	
	arr := []int{}
	for v := range set {
		arr = append(arr, v)
	}
	
	var ans int = len(arr)
	fmt.Println(ans)
}
