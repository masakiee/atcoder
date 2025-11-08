package main

import (
	"fmt"
	"strings"
)

var oddGoodTap = map[string]struct{}{
	"R": {},
	"U": {},
	"D": {},
}

var evenGoodTap = map[string]struct{}{
	"L": {},
	"U": {},
	"D": {},
}

func main() {
	var S string
	fmt.Scanf("%s", &S)

	for i, t := range strings.Split(S, ""){
		tapMap := evenGoodTap
		if (i+1) % 2 == 1 {
			tapMap = oddGoodTap
		}
		_, ok := tapMap[t]
		if !ok {
			fmt.Println("No")
			return
		}
	}
	
	fmt.Println("Yes")
}
