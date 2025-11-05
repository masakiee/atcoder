package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	for i := 1; i <= len(S)-1; i++ {
		if S[i] == S[i-1] {
			fmt.Println("Bad")
			return
		}
	}
	fmt.Println("Good")
}
