package main

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)

	A := make([][]rune, H)
	for i := 0; i < H; i++ {
		A[i] = make([]rune, W)
		var r string
		fmt.Scanf("%s", &r)
		for j, c := range r {
			A[i][j] = c
		}
	}

	for {
		// 1. 行を検査して、順次行削除
		var delRowIndexes []int
		for i := 0; i < len(A); i++ {
			if deletable(row(A, i)) {
				delRowIndexes = append(delRowIndexes, i)
			}
		}
		A = deleteRows(A, delRowIndexes)

		// 2. 列を検査して、順次列削除
		var delColIndexes []int
		for j := 0; j < len(A[0]); j++ {
			if deletable(col(A, j)) {
				delColIndexes = append(delColIndexes, j)
			}
		}
		A = deleteColumns(A, delColIndexes)
		// 3. 消せる行列がなければbreak

		if len(delRowIndexes) == 0 && len(delColIndexes) == 0 {
			break
		}
	}
	
	for _, r := range A {
		fmt.Println(string(r))
	}
}

func row(A[][]rune, i int) []rune {
	return A[i]
}
func col(A[][]rune, j int) []rune {
	var col []rune
	for _, r := range A {
		col = append(col, r[j])
	}
	return col
}
func deletable(A[]rune) bool {
	for _, v := range A {
		if v == '#' {
			return false
		}
	}
	return true
}
func deleteRows(A[][]rune, indexes []int) [][]rune {
	arr := [][]rune{}
	for i, r := range A{
		if !containsInt(indexes, i) {
			arr = append(arr, r)
		}
	}

	return arr
}
func deleteColumns(A[][]rune, indexes []int) [][]rune {
	arr := [][]rune{}
	for _, r := range A {
		var newRow []rune
		for j, c := range r {
			if !containsInt(indexes, j) {
				newRow = append(newRow, c)
			}
		}
		arr = append(arr, newRow)
	}

	return arr
}
func containsInt(arr []int, v int) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}

	return false
}
