package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)


	for {
		if isPrimeNumber(N) {
			fmt.Println(N)
			return
		}
		N++
	}
}

func isPrimeNumber(n int) bool {
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
