package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)

	ans := 0
	balance := 100
	if balance >= X {
		fmt.Println(0)
		return
	}

	for {
		balance += balance / 100
		ans += 1

		if balance >= X {
			break
		}
	}

	fmt.Println(ans)
}
