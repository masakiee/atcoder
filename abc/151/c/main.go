package main

import (
	"fmt"
)

type Result struct {
	AC bool
	WA int
}

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	ret := make(map[int]Result)
	for i := 0; i < M; i++ {
		var p int
		var s string
		fmt.Scanf("%d %s", &p, &s)
		r := ret[p]
		if !r.AC {
			switch s {
			case "AC":
				r.AC = true
			case "WA":
				r.WA++
			}
		}
		ret[p] = r
	}
	
	acNum := 0
	waNum := 0
	for _, r := range ret {
		if r.AC {
			acNum++
			waNum += r.WA
		}
	}
	fmt.Printf("%d %d\n", acNum, waNum)
}
