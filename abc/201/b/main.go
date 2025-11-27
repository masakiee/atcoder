package main

import (
	"fmt"
	"sort"
)

type Mountain struct {
	name string
	height int
}

func main() {
	var N int
	fmt.Scan(&N)
	M := make([]Mountain, N)
	for i := range M {
		m := Mountain{}
		fmt.Scanf("%s %d", &m.name, &m.height)
		M[i] = m
	}

	sort.Slice(M, func (i, j int) bool {
		return M[i].height > M[j].height
	})
	
	ans := M[1].name
	fmt.Println(ans)
}
