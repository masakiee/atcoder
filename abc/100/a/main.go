package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	a := toInt(arr[0])
	b := toInt(arr[1])


	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}


}

func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
