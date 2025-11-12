package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)
	
	min := minInts(A, B)
	var ans int = 1
	var i int = 2
	for i<=min {
		if A % i == 0 && B % i == 0 && isPrimeFactor(i) {
			A /= i
			B /= i
			ans *= i
		} else {
			i++
		}
	}
	ans *= A * B

	fmt.Println(ans)
}

func isPrimeFactor(n int) bool {
	divs := []int{}
	for i:=1; i<=int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			divs = append(divs, n/i)
			if n/i != i {
				divs = append(divs, i)
			}
		}
	}
	return len(divs) == 2
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
