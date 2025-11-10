package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var S string
	fmt.Scan(&S)
	
	i := N / 2
	var ans bool = S[0:i] == S[i:]
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
