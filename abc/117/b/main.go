package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	L := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &L[i])
	}

	sort.Ints(L)
	ok := sumInts(L[0:len(L)-1]...) > L[len(L)-1]
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
