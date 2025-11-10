package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, A, B int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	
	var ans int = 0
	for i:=1; i<=N; i++ {
		sum := digitSum(i)
		if A <= sum && sum <= B {
			ans += i
		}
	}

	fmt.Println(ans)
}

func digitSum(v int) int {
	sum := 0
	s := strconv.Itoa(v)
	for _, d := range strings.Split(s, "") {
		i, _ := strconv.Atoi(d)
		sum += i
	}
	return sum
}
