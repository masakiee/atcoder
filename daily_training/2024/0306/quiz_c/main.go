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
	pp := strings.Split(scanner.Text(), " ")
	num_pp := make([]int, len(pp))
	for i, v := range pp {
		fmt.Sscanf(v, "%d", &num_pp[i])
	}
	cnt := find1(num_pp, num_pp[n-1-1], 1)

	fmt.Printf("%d\n", cnt)
}

func find1(pp []int, pi int, cnt int) int {
	if pi == 1 {
		return cnt
	}
	ppi := pp[pi-1-1]
	return find1(pp, ppi, cnt+1)
}
