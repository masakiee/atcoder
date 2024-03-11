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
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	scanner.Scan()
	s := scanner.Text()
	ss := strings.Split(s, "")
	a := false
	b := false
	c := false
	for i := 0; i < n; i++ {
		if ss[i] == "A" {
			a = true
		} else if ss[i] == "B" {
			b = true
		} else if ss[i] == "C" {
			c = true
		}

		if a && b && c {
			fmt.Printf("%d\n", i+1)
			return
		}
	}
}
