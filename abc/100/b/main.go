package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	d := float64(toInt(arr[0]))
	n := toInt(arr[1])

	if n == 100 {
		n = n + 1
	}
	ret := int(math.Pow(100, d)) * n

	fmt.Println(ret)
}

func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
