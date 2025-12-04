package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	ans := 0
	for i:=1; i<=N; i++ {
		if i % 2 == 0 {
			ans += powInt(i, 3)
		} else {
			ans += -1 * powInt(i, 3)
		}
	}
	fmt.Println(ans)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

