package main

import "fmt"

func main() {
	var letter int
	fmt.Print("Input letter: ")
	fmt.Scan(&letter)

	if letter == letter {
		fmt.Printf("%d adalah palindrom", letter)
	} else {
		fmt.Printf("%d adalah tidak palindrom", letter)
	}
}
