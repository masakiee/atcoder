package main

import (
	"fmt"
	"strings"
)

func main() {
	var H, W, X, Y int
	fmt.Scanf("%d %d %d %d", &H, &W, &X, &Y)
	X--
	Y--
	S := make([]string, H)
	for i := range S {
		fmt.Scan(&S[i])
	}
	
	ans := 1
	SRow := strings.Split(S[X], "")
	SCol := []string{}
	for _, row := range S {
		SCol = append(SCol, string(row[Y]))
	}
	for i:=Y-1; i>=0; i-- {
		if isObstacle(SRow[i]) {
			break
		}
		ans++
	}
	for i:=Y+1; i<W; i++ {
		if isObstacle(SRow[i]) {
			break
		}
		ans++
	}
	for i:=X-1; i>=0; i-- {
		if isObstacle(SCol[i]) {
			break
		}
		ans++
	}
	for i:=X+1; i<H; i++ {
		if isObstacle(SCol[i]) {
			break
		}
		ans++
	}

	fmt.Println(ans)
}

func isObstacle(s string) bool {
	return s == "#"
}
