package main

import (
	"fmt"
	"sort"
)

func main() {
	var P, Q, R int
	fmt.Scanf("%d %d %d", &P, &Q, &R)

	t := []int{P, Q, R}
	sort.Ints(t)
	var ans int = t[0] + t[1]
	fmt.Println(ans)
}
