package main

import (
	"fmt"
)

func main() {
	var N, i int
	fmt.Scanf("%d %d", &N, &i)
	ans := N - i + 1
	fmt.Println(ans)
}
