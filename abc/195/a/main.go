package main

import (
	"fmt"
)

func main() {
	var M, H int
	fmt.Scan(&M)
	fmt.Scan(&H)
	
	if H % M == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
