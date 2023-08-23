package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range intSlice {
		if number%2 == 0 {
			fmt.Printf("%v is Even", number)
			fmt.Println()
		} else {
			fmt.Printf("%v is Odd", number)
			fmt.Println()
		}
	}
}
