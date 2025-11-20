package main

import (
	"fmt"
	"sort"
)

type Worker struct {
	Idx, A, B int
}

func main() {
	var N int
	fmt.Scan(&N)
	W := make([]Worker, N)
	for i := range W {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		W[i] = Worker{i, a, b}
	}

	sort.Slice(W, func (i, j int) bool {
		return W[i].A < W[j].A
	})
	minA1 := W[0]
	minA2 := W[1]
	sort.Slice(W, func (i, j int) bool {
		return W[i].B < W[j].B
	})
	minB1 := W[0]
	minB2 := W[1]

	var ans int
	if minA1.Idx != minB1.Idx {
		ans = maxInts(minA1.A, minB1.B)
	} else {
		t1 := maxInts(minA1.A, minB2.B)
		t2 := maxInts(minA2.A, minB1.B)
		t3 := minA1.A + minB1.B
		ans = minInts(t1, t2, t3)
	}

	fmt.Println(ans)
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
}

func minInts(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
	}
	return min
}
