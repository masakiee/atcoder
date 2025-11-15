package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var X string
	fmt.Scan(&X)

	d := make([]int, len(X))
	for i, v := range strings.Split(X, "") {
		d[i], _ = strconv.Atoi(v)
	}
	sort.Ints(d)
	var dstr string
	for i:=1; i<=9; i++ {
		idx := slices.Index(d, i)
		if idx != -1 {
			dstr = strconv.Itoa(d[idx])
			d = append(d[0:idx], d[idx+1:]...)
			break
		}
	}

	for _, v := range d {
		dstr += strconv.Itoa(v)
	}

	fmt.Println(dstr)
}
