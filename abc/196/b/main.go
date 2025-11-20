package main

import (
	"fmt"
	"strings"
)

func main() {
	var X string
	fmt.Scan(&X)
	
	arr := strings.Split(X, ".")
	var ans string = arr[0]
	
	fmt.Println(ans)
}
