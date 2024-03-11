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
	n := toInt(scanner.Text())
	scanner.Scan()
	arr := make([]int, n)
	for i, ai := range strings.Split(scanner.Text(), " ") {
		arr[i] = toInt(ai)
	}
	min := -1
	for _, a := range arr {
		count := 0
		for a%2 == 0 {
			a /= 2
			count++
		}
		if min == -1 || count < min {
			min = count
		}
		if min == 0 {
			break
		}
	}
	fmt.Println(min)
}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
