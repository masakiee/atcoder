package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	
	for i:=1; i<=9; i++ {
		if N%i == 0 && N/i <= 9 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
