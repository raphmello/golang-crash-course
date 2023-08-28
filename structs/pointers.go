package main

import "fmt"

func main() {
	stringSlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(stringSlice)
	fmt.Println(stringSlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
