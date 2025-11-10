package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	
	var even bool = (a * b) % 2 == 0
	if even {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	} 
}
