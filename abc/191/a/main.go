package main

import (
	"fmt"
)

func main() {
	var V, T, S, D int
	fmt.Scanf("%d %d %d %d", &V, &T, &S, &D)
	
	tx := V * T
	sx := V * S
	if tx <= D && D <= sx {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
