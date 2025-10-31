package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var T, A float64
	fmt.Scanf("%d", &N)
	fmt.Scanf("%f %f", &T, &A)

	idealH := (T - A) / 0.006
	
	H := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%f", &H[i])
	}

	min := struct{
		diff float64
		idx int
	}{
		diff: math.MaxFloat64,
		idx: -1,
	}
	for i, h := range H {
		if abs(h - idealH) < min.diff {
			min.diff = abs(h - idealH)
			min.idx = i + 1
		}
	}
	
	fmt.Println(min.idx)
}

func abs(v float64) float64 {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
