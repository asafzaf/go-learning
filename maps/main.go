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
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Color: %s, Hex: %s\n", color, hex)
	}
}
