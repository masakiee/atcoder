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

}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
