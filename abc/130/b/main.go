package main

import (
	"fmt"
)

func main() {
	var N, X int
	fmt.Scanf("%d %d", &N, &X)

	L := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &L[i])
	}

	di := 0
	for i, li := range L {
		n := i+2
		dn := di + li
		if dn > X {
			fmt.Println(n-1)
			return
		}
		di = dn
	}
	fmt.Println(N+1)
}
