package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	
	ans := S[1:] + S[:1]
	fmt.Println(ans)
}
