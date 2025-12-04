package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	var N, K int
	var S string
	fmt.Scanf("%d %d", &N, &K)
	fmt.Scan(&S)
	
	m := make(map[string]int)
	for i:=0; i<=N-K; i++ {
		substr := S[i:i+K]
		cnt := 0
		for j:=0; j<=N-K; j++ {
			if strings.HasPrefix(S[j:], substr) {
				cnt++
			}
		}
		m[substr] = cnt
	}
	ansMax := 0
	for _, cnt := range m {
		ansMax = maxInts(ansMax, cnt)
	}
	ansStrs := []string{}
	for str, cnt := range m {
		if ansMax == cnt {
			ansStrs = append(ansStrs, str)
		}
	}
	fmt.Println(ansMax)
	sort.Slice(ansStrs, func (i, j int) bool {
		return ansStrs[i] < ansStrs[j]
	})
	for _, v := range ansStrs {
		fmt.Printf("%s ", v)
	}
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
