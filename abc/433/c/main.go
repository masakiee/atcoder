package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part struct {
	Number, StartIdx, EndIdx int
}
func (p Part) Len() int {
	return p.EndIdx - p.StartIdx + 1
}

func main() {
	var S string
	fmt.Scan(&S)
	
	nums := []int{}
	for _, s := range strings.Split(S, "") {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}
	l:=0
	parts := []Part{}
	for l<len(S) {
		r:=l
		for r < len(S) && nums[r] == nums[l] {
			r++
		}
		parts = append(parts, Part{nums[l], l, r-1})
		l=r
	}

	ans := 0
	for i:=0; i<len(parts)-1; i++ {
		lpart, rpart := parts[i], parts[i+1]
		if lpart.Number+1 != rpart.Number {
			continue
		}
		ans += min(lpart.Len(), rpart.Len())
	}

	fmt.Println(ans)
}
