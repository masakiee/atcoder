package main

import (
	"fmt"
)

func main() {
	var r, D, x2000 int
	fmt.Scanf("%d %d %d", &r, &D, &x2000)

	var xi int = x2000
	for i := 0; i < 10; i++ {
		fi := r * xi - D
		fmt.Println(fi)
		xi = fi
	} 
}
