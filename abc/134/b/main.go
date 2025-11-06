package main

import (
	"fmt"
)

func main() {
	var N, D int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &D)
	
	per := 2 * D + 1
	var ans int = (N+per-1)/per
	fmt.Println(ans)
}
