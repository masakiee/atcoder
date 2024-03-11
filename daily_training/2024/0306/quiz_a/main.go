package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()                        // １行分の入力を取得する
	var num int
	fmt.Sscanf(scanner.Text(), "%d", &num)
	if num == 0 {
		fmt.Println("No")
		return
	}
	if num%100 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
