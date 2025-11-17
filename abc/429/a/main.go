package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	for i:=1; i<=N; i++ {
		if i <= M {
			fmt.Println("OK")
		} else {
			fmt.Println("Too Many Requests")
		}
	}
}
