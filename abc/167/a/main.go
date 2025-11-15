package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)
	
	var ans bool = S == T[0:len(T)-1]
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
