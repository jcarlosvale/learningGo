package main

import "fmt"

func main() {
	//colors := map[string]string{
	//	"red":   "#ff000",
	//	"green": "#4bf5443",
	//}

	//var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffff"

	//delete(colors, "white")
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
