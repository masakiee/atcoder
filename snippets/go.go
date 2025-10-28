package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ans int
	out(ans)
}

var sc = bufio.NewScanner(os.Stdin)

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := range N {
		ret[i] = getI()
	}
	return ret
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func out(x ...interface{}) {
	fmt.Println(x...)
}
