package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Scanf("%d", &K)
	
	var oddnum int
	for i := 1; i <= K; i++ {
		if i % 2 == 1 {
			oddnum++
		}
	}
	evennum := K - oddnum

	fmt.Println(oddnum * evennum)
}
