package main

import (
	"fmt"
)

func main() {
	var N, S, D int
	fmt.Scanf("%d %d %d", &N, &S, &D)

	var ans string = "No"
	for i:=0; i<N; i++ {
		var time, attack int
		fmt.Scanf("%d %d", &time, &attack)

		if time < S && attack > D && ans == "No" {
			ans = "Yes"
		}
	}
	
	fmt.Println(ans)
}
