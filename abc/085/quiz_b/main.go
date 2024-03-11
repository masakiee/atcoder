package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dmap := map[int]bool{}
	for scanner.Scan() {
		di := toInt(scanner.Text())
		if _, found := dmap[di]; !found {
			dmap[di] = true
		}
	}

	fmt.Println(len(dmap))
}
func toInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
