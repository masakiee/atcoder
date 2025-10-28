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
	firstline := strings.Split(scanner.Text(), " ")
	N := toInt(firstline[0])
	C := toInt(firstline[1])
	scanner.Scan()
	A := make([]int, N)
	for i, ai := range strings.Split(scanner.Text(), " ") {
		A[i] = toInt(ai)
	}

}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
