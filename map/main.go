package main

import "fmt"

func main() {
	colors := map[string]string{
		"R": "red",
		"B": "blue",
		"G": "green",
	}

	colors["Y"] = "yellow"

	delete(colors, "Y")

	fmt.Println(colors["Y"])

	printMap(colors)
}

func printMap(c map[string]string) {
	for colorCode, color := range c {
		fmt.Printf("%v is the code for %v\n", colorCode, color)
	}
}
