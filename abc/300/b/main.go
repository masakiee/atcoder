package main

import (
	"fmt"
	"strings"
)

func main() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	A := getGrid(H, W)
	B := getGrid(H, W)

	if solve(A, B) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve(A, B [][]string) bool {
	for i := 0; i < h(A); i++ {
		for j := 0; j < w(A); j++ {
			A_ := shiftGrid(A, i, j)
			if isSame(A_, B){
				return true
			}
		}
	}
	return false
}

func h(grid[][]string) int {
	return len(grid)
}
func w(grid[][]string) int {
	return len(grid[0])
}

func shiftGrid(grid [][]string, row, col int) [][]string {
	ret := make([][]string, h(grid))
	for i := range grid {
		ret[i] = make([]string, w(grid))
		for j := range grid[i] {
			ret[i][j] = grid[(i+row)%h(grid)][(j+col)%w(grid)]
		}
	}
	return ret
}


func getGrid(H, W int) [][]string {
	A := make([][]string, H)
	for i := 0; i < H; i++ {
		A[i] = make([]string, W)
		var row string
		fmt.Scan(&row)
		for j, c := range strings.Split(row, "") {
			A[i][j] = c
		}
	}
	return A
}

func isSame(a, b [][]string) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func sumInts(arr ...int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}

func minInts(arr ...int) int {
	min := arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func filter[T any](arr []T, f func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
