package main

import (
	"fmt"
)

func main() {
	var S int
	fmt.Scanf("%d", &S)

	done := make(map[int]struct{})
	done[S] = struct{}{}
	prev := S
	for i := 2; i <= 1000000; i++ {
		ai := f(prev)
		_, exists := done[ai]
		if exists {
			fmt.Println(i)
			return
		}
		done[ai] = struct{}{}
		prev = ai
	}
}

func f(n int) int {
	if n % 2 == 0 {
		return n / 2
	} else {
		return 3 * n + 1
	}
}
