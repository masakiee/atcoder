package main

import (
	"fmt"
)

var forecast = map[string]string{
	"Sunny": "Cloudy",
	"Cloudy": "Rainy",
	"Rainy": "Sunny",
}
func main() {
	var S string
	fmt.Scanf("%s", &S)
	
	var ans string = forecast[S]
	fmt.Println(ans)
}
