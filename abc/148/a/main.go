package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)

	var ans int = 6 / (A * B)
	fmt.Println(ans)
}
