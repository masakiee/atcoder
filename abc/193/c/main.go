package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	cnt := make(map[int]struct{})
	for a:=2; a*a<=N; a++ {
		x := a*a
		for x <= N {
			cnt[x] = struct{}{}
			x *= a
		}
	}
	ans := N - len(cnt)
	fmt.Println(ans)
}
