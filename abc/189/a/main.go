package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	cc := strings.Split(S, "")
	c1, c2, c3 := cc[0], cc[1], cc[2]
	if c1 == c2 && c2 == c3 {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}
