package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	k, _ := strconv.Atoi(inputs[1])
	
	ans := ((n - 1) + (k - 1) - 1) / (k - 1)
	fmt.Println(ans)
}

