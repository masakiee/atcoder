package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)
	
	var ans string = T + S
	fmt.Println(ans)
}
