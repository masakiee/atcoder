package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &P[i])
	}

	ans := 0
	for i:=1; i<=N-2; i++ {
		arr := []int{P[i-1], P[i], P[i+1]}
		sort.Ints(arr)
		if arr[1] == P[i] {
			ans++
		}
	}
	
	fmt.Println(ans)
}
