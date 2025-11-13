package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colors)

	printMap(colors)

	colors["yellow"] = "#ffff00"
	printMap(colors)
	delete(colors, "green")
	printMap(colors)

	value, ok := colors["purple"]
	if ok {
		fmt.Println("The hex code for purple is:", value)
	} else {
		fmt.Println("Purple not found in the map")
	}

	value2, ok := colors["yellow"]
	if ok {
		fmt.Println("The hex code for yellow is:", value2)
	} else {
		fmt.Println("Yellow not found in the map")
	}
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Color: %s, Hex: %s\n", color, hex)
	}
}
