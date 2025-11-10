package main

import (
	"fmt"
)

type Step struct {
	t, x, y int
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	steps := make([]Step, N+1)
	steps[0] = Step{0, 0, 0}
	for i := 1; i <= N; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)
		steps[i] = Step{t, x, y}
	}
	
	for i := 1; i <= N; i++ {
		dt := steps[i].t - steps[i-1].t
		dx := steps[i].x - steps[i-1].x
		dy := steps[i].y - steps[i-1].y
		absD := absInt(dx) + absInt(dy)
		if dt / absD >= 1 && (absD - dt) % 2 == 0 {
			continue
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func absInt(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}
