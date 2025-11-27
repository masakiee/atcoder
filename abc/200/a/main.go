package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var ans int = (N + 99) / 100
	fmt.Println(ans)
}
