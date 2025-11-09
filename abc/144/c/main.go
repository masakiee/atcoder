package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	divs := [][]int{}
	nSqrt := int(math.Sqrt(float64(N)))
	for i:=nSqrt; i>=1; i-- {
		if N%i == 0 {
			divs = append(divs, []int{N/i, i})
			ans := (N/i-1) + (i-1)
			fmt.Println(ans)
			return
		}
	} 
}
