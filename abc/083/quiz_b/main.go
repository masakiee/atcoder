package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	N := toInt(line[0])
	A := toInt(line[1])
	B := toInt(line[2])
	sum := 0
	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		digitSum := 0
		for _, c := range s {
			digitSum += int(c - '0')
		}
		if A <= digitSum && digitSum <= B {
			sum += i
		}
	}
	fmt.Println(sum)
}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
