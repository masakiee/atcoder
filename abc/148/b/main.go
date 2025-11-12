package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var S, T string
	fmt.Scanf("%s %s", &S, &T)

	var ans string
	for i:=0; i<N; i++ {
		ans += string(S[i]) + string(T[i])
	}
	
	fmt.Println(ans)
}
