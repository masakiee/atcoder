package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N := toInt(scanner.Text())

	cfz := 0
	for N%2 == 0 {
		N = N / 2
		cfz++
	}

	fmt.Println(cfz)
}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
