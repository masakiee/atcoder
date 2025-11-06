package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	
	nstr := strconv.Itoa(N)
	var ans int
	for i := 1; i <= len(nstr); i++ {
		if i % 2 == 0 {
			continue
		}
		if i < len(nstr) {
			ans += int(math.Pow10(i)) - int(math.Pow10(i-1))
		} else {
			ans += N - int(math.Pow10(i-1)) + 1
		}
	}
	fmt.Println(ans)
}
