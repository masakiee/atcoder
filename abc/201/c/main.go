package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	used := []int{}
	cannot := []int{}
	for i, c := range strings.Split(S, "") {
		switch c {
		case "o":
			used = append(used, i)
		case "x":
			cannot = append(cannot, i)
		}
	}
	ans := 0
	for i:=0; i<=9999; i++ {
		str := fmt.Sprintf("%04d", i)
		ok := true
		for _, v := range used {
			if !strings.Contains(str, strconv.Itoa(v)) {
				ok = false
				break
			}
		}
		for _, v := range cannot {
			if strings.Contains(str, strconv.Itoa(v)) {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}

	fmt.Println(ans)
}
