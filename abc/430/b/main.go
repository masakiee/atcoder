package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	
	S := make([][]rune, N)
	for i := 0; i < N; i++ {
		var l string
		fmt.Scanf("%s", &l)
		S[i] = make([]rune, N)
		for j, c := range l {
			S[i][j] = c
		}
	}

	ret := make(map[string]struct{})
	for i := 0; i < N-M+1; i++ {
		for j := 0; j < N-M+1; j++ {
			str := blockToString(S, i, j, M)
			ret[str] = struct{}{}
		}
	}
	
	ans := len(ret)
	fmt.Println(ans)
}

func blockToString(arr [][]rune, i, j, m int) string {
	var ret string
	for row := i; row < i+m; row++ {
		for col := j; col < j+m; col++ {
			ret += string(arr[row][col])
		}
	}
	return ret
}
