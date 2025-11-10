package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)
	
	if traverse(S, 0, len(S)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func traverse(s string, pos int, sLen int) bool {
	if sLen == pos {
		return true
	}

	var words = []string{
		"dreamer", "dream",
		"eraser", "erase",
	}

	for _, word := range words {
		ok := len(s) >= len(word) && s[0:len(word)] == word && traverse(s[len(word):], pos+len(word), sLen)
		if ok {
			return true
		}
	}

	return false
}
