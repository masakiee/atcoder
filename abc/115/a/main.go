package main

import (
	"fmt"
)

func main() {
	var D int
	fmt.Scanf("%d", &D)

	fmt.Print("Christmas")
	for i := D; i < 25; i++ {
		fmt.Print(" Eve")
	}

	fmt.Print("\n")
}
