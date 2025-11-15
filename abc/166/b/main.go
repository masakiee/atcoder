package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	okashi := map[int]struct{}{}
	for i := 0; i < K; i++ {
		var d int
		fmt.Scan(&d)
		for j := 0; j < d; j++ {
			var Aij int
			fmt.Scanf("%d", &Aij)
			okashi[Aij] = struct{}{}
		}
	}
	
	var ans int = N - len(okashi)
	fmt.Println(ans)
}
