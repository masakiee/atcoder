package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	numOdd := (N+2-1)/2

	var ans float64 = float64(numOdd) / float64(N)
	fmt.Println(ans)
}
