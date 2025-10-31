package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	W := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&W[i])
	}

	if !isUnique(W) {
		fmt.Println("No")
		return
	}

	var c byte
	for i, w := range W {
		if i != 0 && c != w[0] {
			fmt.Println("No")
			return
		}
		c = w[len(w)-1]
	}
	
	fmt.Println("Yes")
}

func isUnique(arr []string) bool {
	m := map[string]bool{}
	for _, v := range arr {
		if m[v] {
			return false
		}
		m[v] = true
	}

	return true
}
