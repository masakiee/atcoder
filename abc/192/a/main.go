package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)
	
	var ans int = 100 - X%100
	fmt.Println(ans)
}
