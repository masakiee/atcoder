package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	digitN := numDigits(N)
	var ans int
	for i:=1; i<=digitN/2; i++ {
		if i*2 == digitN {
			first, last := splitN(N)
			num := first - (pow(10, i-1) - 1)
			if first > last {
				num--
			}
			ans += num
		} else {
			ans += 9 * pow(10, i-1)
		}
		
	}

	fmt.Println(ans)
}

func splitN(n int) (int, int) {
	s := strconv.Itoa(n)

	first, err := strconv.Atoi(s[0:len(s)/2])
	if err != nil {
		panic(err)
	}
	last, err := strconv.Atoi(s[len(s)/2:])
	if err != nil {
		panic(err)
	}
	return first, last
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func numDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}
