package main

import (
	"fmt"
)

func main() {
	var A, B float64
	fmt.Scan(&A)
	fmt.Scan(&B)

	ans := (A - B) / A * 100
	fmt.Printf("%.3f\n", ans)
}
