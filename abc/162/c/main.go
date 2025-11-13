package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Scan(&K)

	memoMap := make(map[int]map[int]int)
	tmp := []int{}
	for a:=1; a<=K; a++ {
		for b:=1; b<=K; b++ {
			k1, k2 := a, b
			if a < b {
				k1, k2 = b, a
			}
			m, exists := memoMap[k1][k2]
			var value int
			if exists {
				value = m
			} else {
				value = gcd(k1, k2)
				if memoMap[k1] == nil {
					memoMap[k1] = make(map[int]int)
				}
				memoMap[k1][k2] = value
			}
			tmp = append(tmp, value)
		}
	}

	arr := []int{}
	for c:=1; c<=K; c++ {
		for _, ab := range tmp {
			k1, k2 := c, ab
			if c < ab {
				k1, k2 = ab, c
			}
			m, exists := memoMap[k1][k2]
			var value int
			if exists {
				value = m
			} else {
				value = gcd(k1, k2)
				if memoMap[k1] == nil {
					memoMap[k1] = make(map[int]int)
				}
				memoMap[k1][k2] = value
			}
			arr = append(arr, gcd(ab, c))
		}
	}
	
	var ans int = sumInts(arr...)
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		na, nb := b, a%b
		return gcd(na, nb)
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
