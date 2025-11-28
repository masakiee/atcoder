package main

import (
	"fmt"
	"math"
	"sort"
)

type Player struct {
	idx, rate int
}
func main() {
	var N int
	fmt.Scan(&N)
	size := pow(2, N)
	players := make([]Player, size)
	for i:=0; i<size; i++ {
		p := Player{}
		p.idx = i+1
		fmt.Scanf("%d", &p.rate)
		players[i] = p
	}
	players1 := players[0:size/2]
	players2 := players[size/2:]

	sort.Slice(players1, func(i, j int) bool {
		return players1[i].rate > players1[j].rate
	})
	sort.Slice(players2, func(i, j int) bool {
		return players2[i].rate > players2[j].rate
	})

	var semiFinalist Player
	if players1[0].rate > players2[0].rate {
		semiFinalist = players2[0]
	} else {
		semiFinalist = players1[0]
	}

	var ans int = semiFinalist.idx
	fmt.Println(ans)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
