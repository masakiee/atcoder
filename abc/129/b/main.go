package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	W := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &W[i])
	}

	ans := math.MaxInt
	for i:=1; i<N; i++ {
		s1 := sumInts(W[0:i+1]...)
		s2 := sumInts(W[i+1:]...)
		if ans > absInt(s2 - s1) {
			ans = absInt(s2 - s1)
		}
	}
	
	fmt.Println(ans)
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
