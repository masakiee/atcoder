package main

import (
	"fmt"
)

func main() {
	var H, B int
	fmt.Scanf("%d", &H)
	fmt.Scanf("%d", &B)
	
	var ans int = maxInts(0, H - B)
	fmt.Println(ans)
}

func maxInts(arr ...int) int {
	max := arr[0]
	for _, v := range arr[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
