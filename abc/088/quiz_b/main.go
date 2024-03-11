package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := toInt(scanner.Text())
	scanner.Scan()
	arr := make([]int, n)
	for i, ai := range strings.Split(scanner.Text(), " ") {
		arr[i] = toInt(ai)
	}
	// 配列をクイックソートする(昇順)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	score := 0
	for i, v := range arr {
		if i%2 == 0 {
			score += v
		} else {
			score -= v
		}
	}

	fmt.Println(score)
}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
