package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &P[i])
	}

	cnt := 0
	for i := 1; i <= N; i++ {
		v := P[i-1]
		if v != i {
			cnt++
		}
	}
	
	ans := "NO"
	if cnt == 2 || cnt == 0 {
		ans = "YES"
	}
	fmt.Println(ans)
}
