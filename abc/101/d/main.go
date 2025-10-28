package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	K := getI()

	ret := []int{}
	n := 1
	for {
		if getNSn(n) <= getNSn(K) {
			ret = append(ret, n)
			if len(ret) >= K {
				break
			}
		}
		n++
		if n > K {
			break
		}
	}

	for _, r := range ret {
		out(r)
	}
}

func getI() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func getSn(n int) int {
	ss := strings.Split(strconv.Itoa(n), "")
	var sum int
	for _, si_ := range ss {
		si, _ := strconv.Atoi(si_)
		sum += si
	}

	return sum
}

func getNSn(n int) float64 {
	return float64(n) / float64(getSn(n))
}

var sc = bufio.NewScanner(os.Stdin)

func out(x ...interface{}) {
	fmt.Println(x...)
}
