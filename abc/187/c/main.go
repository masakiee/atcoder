package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	memo := make(map[string]bool)
	var str string
	for i:=0; i<N; i++ {
		fmt.Scanf("%s", &str)
		memo[str] = true
	}

	for s := range memo {
		if memo["!"+s] {
			fmt.Println(s)
			return
		}
	}
	fmt.Println("satisfiable")
}
