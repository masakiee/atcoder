package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scanf("%d", &X)

	if X == 3 || X == 5 || X == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
