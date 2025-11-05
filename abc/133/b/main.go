package main

import (
	"fmt"
)

func main() {
	var N, D int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &D)

	X := make([][]int, N)
	for i := 0; i < N; i++ {
		X[i] = make([]int, D)
		for j := 0; j < D; j++ {
			fmt.Scanf("%d", &X[i][j])
		}
	}

	squareNumbers := make(map[int]struct{})
	maxSquareNumber := (20-(-20)) * (20-(-20)) * 10
	for i := 1; i*i <= maxSquareNumber; i++ {
		squareNumbers[i*i] = struct{}{}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			_, ok := squareNumbers[squareDist(X[i], X[j])]
			if ok {
				ans++
			}
		}
	}
	
	fmt.Println(ans)
}

func squareDist(a, b []int) int {
	dist := 0
	for i := range a {
		dist += (a[i]-b[i]) * (a[i]-b[i])
	}
	return dist
}
