package main

import (
	"fmt"
)

func main() {
	var X, Y int
	fmt.Scan(&X)
	fmt.Scan(&Y)

	diff := absInt(X - Y)
	if diff < 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
