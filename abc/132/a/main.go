package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	ret := make(map[rune]int)
	for _, c := range S {
		ret[c]++
	}

	ans := len(ret) == 2
	for _, v := range ret {
		if v != 2 {
			ans = false
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
