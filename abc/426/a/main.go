package main

import (
	"fmt"
)

func main() {
	var X, Y string
	fmt.Scanf("%s %s", &X, &Y)

	m := map[string][]string{
		"Ocelot": {"Ocelot", "Serval", "Lynx"},
		"Serval": {"Serval", "Lynx"},
		"Lynx": {"Lynx"},
	}
	
	ms := m[Y]
	if contains(ms, X) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func contains(arr []string, needle string) bool {
	for _, v := range arr {
		if v == needle {
			return true
		}
	}
	return false
}
