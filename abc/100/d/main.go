package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func run(in io.Reader, out io.Writer) {
	var n, m, ans int
	fmt.Fscan(in, &n, &m)
	a := make([]struct{ x, y, z int }, n)
	for i := range a {
		fmt.Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}

	fmt.Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

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
