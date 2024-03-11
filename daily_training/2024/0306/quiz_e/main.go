package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	scanner.Scan()
	pn := make([]int, n)
	for i, v := range strings.Split(scanner.Text(), " ") {
		fmt.Sscanf(v, "%d", &pn[i])
	}

	// k-1, k, k+1 回回転させた時に嬉しい人数のマッピング
	ck := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			ck[(i-pn[i]-1+j+n)%n] += 1
		}
	}
	max := 0
	for _, c := range ck {
		max = calcMax(max, c)
	}
	fmt.Printf("%d\n", max)
}

func calcMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
