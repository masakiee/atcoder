package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	
	fmt.Println(solve(N, K))
}

func solve(N, K int) int {
	var a int = N
	for i:=0; i<K; i++ {
		a = f(a)
	}
	return a
}

var memo = make(map[int]int)

func f(x int) int {
	if cache, ok := memo[x]; ok {
		return cache
	}

	arr := strings.Split(strconv.Itoa(x), "")
	digits := make([]int, len(arr))
	for i, d := range arr {
		digits[i], _ = strconv.Atoi(d)
	}

	sort.Ints(digits)
	g1, g2 := 0, 0
	for i:=0; i<len(digits); i++ {
		pow := int(math.Pow(10.0, float64(len(digits)-1-i)))
		g1 += digits[len(digits)-1-i] * pow
		g2 += digits[i] * pow
	}

	memo[x] = g1 - g2
	return g1 - g2
}
