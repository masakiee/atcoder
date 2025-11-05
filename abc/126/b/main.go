package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)
	
	headCanBeMonth := canBeMonth(S[0:2])
	tailCanBeMonth := canBeMonth(S[2:4])

	if (headCanBeMonth && tailCanBeMonth) {
		fmt.Println("AMBIGUOUS")
	} else if (headCanBeMonth) {
		fmt.Println("MMYY")
	} else if (tailCanBeMonth) {
		fmt.Println("YYMM")
	} else {
		fmt.Println("NA")
	}
}

func canBeMonth(s string) bool {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return 1 <= v && v <= 12
}
