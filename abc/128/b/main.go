package main

import (
	"fmt"
	"sort"
)

type Restaurant struct {
	Name string
	Point int
	Idx int
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	restaurants := make([]Restaurant, N)

	for i:=0; i<N; i++ {
		var s string
		var p int
		fmt.Scanf("%s %d", &s, &p)
		restaurants[i] = Restaurant{
			Name: s,
			Point: p,
			Idx: i+1,
		}
	}

	sort.Slice(restaurants, func (i, j int) bool {
		ri := restaurants[i]
		rj := restaurants[j]
		if  ri.Name < rj.Name {
			return true
		} else if ri.Name == rj.Name && ri.Point > rj.Point {
			return true
		} else {
			return false
		}
		
	})
	
	for _, r := range restaurants {
		fmt.Println(r.Idx)
	}
}
