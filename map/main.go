package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	var colors2 map[string]string

	colors3 := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "red")

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
