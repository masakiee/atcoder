package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &Y)

	max10k := minInts(N, Y/10000)
	for i:=0; i<=max10k; i++ {
		rest := Y - 10_000 * i
		max5k := minInts(N-i, rest / 5_000)
		for j:=0; j<=max5k; j++ {
			num1k := (rest - 5_000 * j) / 1_000
			if num1k == N-i-j {
				fmt.Printf("%d %d %d\n", i, j, num1k)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}

func minInts(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
	}
	return min
}
