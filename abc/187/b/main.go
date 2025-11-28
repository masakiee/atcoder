package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}
func main() {
	var N int
	fmt.Scan(&N)
	P := make([]Point, N)
	for i := range P {
		p := Point{}
		fmt.Scanf("%f %f", &p.x, &p.y)
		P[i] = p
	}

	ans := 0
	for i:=0; i<N; i++ {
		for j:=i+1; j<N; j++ {
			a := (P[j].y-P[i].y) / (P[j].x-P[i].x)
			if -1.0 <= a && a <= 1.0 {
				ans++
			}
		}
	}
	
	fmt.Println(ans)
}
