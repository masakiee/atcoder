package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	
	A := []int{1}
	for i:=1; i<=N; i++ {
		Ai := 0
		for _, a := range A {
			Ai += f(a)
		}
		A = append(A, Ai)
	}
	
	var ans int = A[len(A)-1]
	fmt.Println(ans)
}

func f(i int) int {
	str := strconv.Itoa(i)
	sum := 0
	for _, c := range strings.Split(str, "") {
		v, _ := strconv.Atoi(c)
		sum += v
	}
	return sum
}
