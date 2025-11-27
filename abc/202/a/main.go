package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	
	var ans int = 21 - a - b - c
	fmt.Println(ans)
}
