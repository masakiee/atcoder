package main

import (
	"fmt"
)

func main() {
	var b string
	fmt.Scanf("%s", &b)
	fmt.Println(basePair(b))
}

func basePair(b string) string {
	var bp string
	switch b {
	case "A":
		bp = "T"
	case "T":
		bp = "A"
	case "C":
		bp = "G"
	case "G":
		bp = "C"
	}

	return bp
}
