package main

import (
	"fmt"
	"math"
)

type Pos struct {
	X, Y float64
}
func main() {
	var N int
	fmt.Scanf("%d", &N)
	P := make([]Pos, N)
	for i := 0; i < N; i++ {
		var x, y float64
		fmt.Scanf("%f %f", &x, &y)
		P[i] = Pos{X: x, Y: y}
	}

	count := float64(numPer(N-1))
	sum := float64(0)
	for i:=0; i<N-1; i++ {
		for j:=i+1; j<N; j++ {
			sum += distance(P[i], P[j]) * count * float64(2)
		}
	}

	var ans float64 = sum / float64(numPer(N))
	fmt.Println(ans)
}

func distance(a, b Pos) float64 {
	return math.Pow((b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y), 0.5)
}

func numPer(n int) int {
	ans := 1
	for i:=n; i>0; i-- {
		ans *= i
	}
	return ans
}
