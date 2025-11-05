package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	var S string
	fmt.Scanf("%d %d", &N, &K)
	fmt.Scan(&S)

	ret := make([]string, len(S))
	for i, c := range strings.Split(S, "") {
		if i == K-1 {
			ret = append(ret, strings.ToLower(c))
		} else {
			ret = append(ret, c)
		}
	}
	
	ans := strings.Join(ret, "")
	fmt.Println(ans)
}
