package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	answer := map[string]int{
		"SUN": 7,
		"MON": 6,
		"TUE": 5,
		"WED": 4,
		"THU": 3,
		"FRI": 2,
		"SAT": 1,
	}
	
	var ans int = answer[S]
	fmt.Println(ans)
}
