package main

import (
	"fmt"
)

func main() {
	var A1, A2, A3 int
	fmt.Scan(&A1)
	fmt.Scan(&A2)
	fmt.Scan(&A3)

	sum := sumInts(A1, A2, A3)
	if sum >= 22 {
		fmt.Println("bust")
	} else {
		fmt.Println("win")
	}
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
