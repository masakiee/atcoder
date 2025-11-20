package main

import (
	"fmt"
	"sort"
)

type Machine struct {
	a, p, x int
}

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]Machine, N)
	for i := range arr {
		m := Machine{}
		fmt.Scanf("%d %d %d", &m.a, &m.p, &m.x)
		arr[i] = m
	}

	arr = filter(arr, func (m Machine) bool {
		return m.x - m.a >= 1
	})
	if len(arr) == 0 {
		fmt.Println(-1)
		return
	}

	sort.Slice(arr, func (i, j int) bool {
		return arr[i].p < arr[j].p
	})

	var ans int = arr[0].p
	fmt.Println(ans)
}

func filter[T any](arr []T, f func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
