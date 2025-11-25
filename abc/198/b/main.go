package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	
	s := strconv.Itoa(N)
	max := len(s)
	for i:=0; i<max; i++ {
		if isPalindrome(s) {
			fmt.Println("Yes")
			return
		}
		s = "0" + s
	}
	fmt.Println("No")
}

func isPalindrome(s string) bool {
	arr := strings.Split(s, "")
	for i:=0; i<len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}
