package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	maxPersons := make([]int, 5)
	for i:=0; i<len(maxPersons); i++ {
		fmt.Scanf("%d", &maxPersons[i])
	}

	minPerson := minInts(maxPersons)
	ans := 5 + (numOfGroup(N, minPerson) - 1)
	fmt.Println(ans)
}

func minInts(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if m > v {
			m = v
		}
	}
	return m
}

func numOfGroup(n, unit int) int {
	return int(math.Ceil(float64(n) / float64(unit)))
}
