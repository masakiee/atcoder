package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	ans := make([]int, N)
	for i := 2; i <= N; i++ {
		var a int
		fmt.Scanf("%d", &a)
		ans[a-1]++
	}
	
	for _, v := range ans {
		fmt.Println(v)
	}
}
