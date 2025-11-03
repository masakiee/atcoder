package main

import (
	"fmt"
)

func main() {
	p := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &p[i])
	}
	var k int
	fmt.Scanf("%d", &k)
	
	for i := 0; i < 5; i++ {
		for j := i+1; j < 5; j++ {
			d := abs(p[i] - p[j])
			if d > k {
				fmt.Println(":(")
				return
			}
		}
	}

	fmt.Println("Yay!")
}

func abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
