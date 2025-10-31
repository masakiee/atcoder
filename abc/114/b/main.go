package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	diffs := make([]int, len(S)-3+1)
	for i := 0; i <= len(S) - 3; i++ {
		s := S[i:i+3]
		diffs[i] = abs(753 - toInt(s))
	}

	sort.Ints(diffs)
	fmt.Println(diffs[0])
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
