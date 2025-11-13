package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	all := sigma(N)
	x3 := 3 * sigma(N/3)
	x5 := 5 * sigma(N/5)
	x15 := 15 * sigma(N/15)
	var ans int = all - x3 - x5 + x15
	fmt.Println(ans)
}

func sigma(n int) int {
	return n * (n+1) / 2
}
