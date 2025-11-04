package main

import (
	"fmt"
	"sort"
)

func main() {
	A := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sort.Slice(A, func(i, j int) bool {
		return (A[i]-1)%10 > (A[j]-1)%10
	})
	ans := 0
	for i, x := range A {
		if i == len(A) - 1 {
			ans += x
		} else if x % 10 == 0 {
			ans += x
		} else {
			ans += (x/10+1)*10
		}
	}
	fmt.Println(ans)
}
