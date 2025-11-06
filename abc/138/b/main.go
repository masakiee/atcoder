package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	
	Rs := []float64{}
	for _, v := range A {
		Rs = append(Rs, float64(1)/float64(v))
	}

	var ans float64 = float64(1) / sumFloats(Rs...)
	fmt.Println(ans)
}

func sumFloats(arr ...float64) float64 {
	var ret float64
	for _, v := range arr {
		ret += v
	}
	return ret
}
