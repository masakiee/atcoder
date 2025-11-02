package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)
	
	r, _ := regexp.Compile("[^ACGT]+")
	ps := r.Split(S, -1)
	size := make([]int, len(ps))
	for i, p := range ps {
		size[i] = len(p)
	}
	fmt.Println(max(size))
}

func max(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}
