package main

import (
	"fmt"
)

type Part struct {
	W int
	H int
	B int
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	W := make([]int, N)
	H := make([]int, N)
	B := make([]int, N)
	parts := []Part{}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &W[i])
		fmt.Scanf("%d", &H[i])
		fmt.Scanf("%d", &B[i])
		parts = append(parts, Part{
			W: W[i],
			H: H[i],
			B: B[i],
		})
	}

	var max int = 0
	for i:=0; i<N; i++ {
		p := []Part{}
		p = append(p, parts[i:]...)
		p = append(p, parts[0:i]...)
		// 頭は1つもパーツをつけなくても良い
		for k:=0; k<=N-1; k++ {
			wHead := 0
			wBody := 0
			happiness := 0
			for _, part := range p[0:k] {
				happiness += part.H
				wHead += part.W
			}
			for _, part := range p[k:] {
				happiness += part.B
				wBody += part.W
			}

			if wHead <= wBody && max < happiness {
				max = happiness
			}
		}
	}
	fmt.Println(max)
}
