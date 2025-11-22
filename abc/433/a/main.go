package main

import (
	"fmt"
)

func main() {
	var X, Y, Z int
	fmt.Scan(&X)
	fmt.Scan(&Y)
	fmt.Scan(&Z)
	
	if (X - Z*Y) >= 0 && (X - Z*Y) % (Z-1) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
