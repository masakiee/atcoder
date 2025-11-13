package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	S := make(map[string]struct{})
	for i := 0; i < N; i++ {
		var s string
		fmt.Scanf("%s", &s)
		S[s] = struct{}{}
	}
	
	var ans int = len(S)
	fmt.Println(ans)
}
