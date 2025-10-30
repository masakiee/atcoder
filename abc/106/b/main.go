package main

import (
	"fmt"
	"math"
)

// 約数の数を返す
func divisors(n int) []int {
	ret := []int{}
	// i が約数の時、n/i も約数になる
	for i := 1; i <= int(math.Sqrt(float64(n))) + 1; i++ {
		if n % i == 0 {
			ret = append(ret, i)
			if n/i != i {
				ret = append(ret, n/i)
			}
		}
	}

	return ret
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	arr := []int{}
	for n := 1; n <= N; n+=2 {
		if  len(divisors(n)) == 8 {
			arr = append(arr, n)
		}
	}
	
	fmt.Println(len(arr))
}
