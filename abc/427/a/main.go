package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	ans := S[:len(S)/2] + S[len(S)/2+1:]
	fmt.Println(ans)
}
