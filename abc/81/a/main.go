package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)
	var ans int
	for _, s := range strings.Split(S, "") {
		if s == "1" {
			ans++
		}
	}
	fmt.Println(ans)
}
