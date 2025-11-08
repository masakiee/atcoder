package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scanf("%d", &N)
	fmt.Scanf("%s", &S)
	num := len(S)

	for i := 0; i < N-1; i++ {
		if S[i] == S[i+1] {
			num--
		}
	}
	
	var ans int = num
	fmt.Println(ans)
}
