package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ss := strings.Split(scanner.Text(), " ")
	// goal = ゴール座標
	// kabe = 壁の座標
	// hummer = ハンマーの座標
	var goal, kabe, hummer int
	fmt.Sscanf(ss[0], "%d", &goal)
	fmt.Sscanf(ss[1], "%d", &kabe)
	fmt.Sscanf(ss[2], "%d", &hummer)

	// 壁の反対側にゴールとハンマーがある場合はゴールできない
	if goal > 0 && kabe > 0 && goal > kabe && hummer > kabe {
		fmt.Println("-1")
		return
	}
	if goal < 0 && kabe < 0 && goal < kabe && hummer < kabe {
		fmt.Println("-1")
		return
	}

	// 原点->壁→ゴールとなっているが、壁の手前側にハンマーがある場合はゴールできる
	if goal > 0 && kabe > 0 && goal > kabe && kabe > hummer {
		var distance int
		if hummer > 0 {
			distance = goal
		} else {
			distance = goal + hummer*-2
		}
		fmt.Printf("%d\n", distance)
		return
	}
	if goal < 0 && kabe < 0 && goal < kabe && kabe < hummer {
		var distance int
		if hummer < 0 {
			distance = goal * -1
		} else {
			distance = goal*-1 + hummer*2
		}
		fmt.Printf("%d\n", distance)
		return
	}

	// それ以外の場合は普通にゴールできる
	fmt.Printf("%d\n", int(math.Abs(float64(goal))))
}
