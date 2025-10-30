package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)

	dx21 := x2 - x1
	dy21 := y2 - y1

	x3 := x2 + dy21 * (-1)
	y3 := y2 + dx21

	x4 := x3 + dx21 * (-1)
	y4 := y3 + dy21 * (-1)
	fmt.Printf("%d %d %d %d\n", x3, y3, x4, y4)
}
