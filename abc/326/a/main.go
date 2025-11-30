package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	
	fmt.Println(solve(x, y))
}

func solve(x, y int) string {
	if y - x > 2 {
		return "No"
	} else if y - x >= -3 {
		return "Yes" 
	} else {
		return "No"
	}
}
