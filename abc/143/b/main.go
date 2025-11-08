package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	D := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &D[i])
	}

	var ans int
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			if i != j {
				ans += D[i] * D[j]
			}
		} 
	}
	
	fmt.Println(ans)
}
