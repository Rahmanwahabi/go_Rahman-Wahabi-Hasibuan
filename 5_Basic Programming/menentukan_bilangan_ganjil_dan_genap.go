package main

import "fmt"

func main() {
	var number int
	fmt.Print("Input Number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("%d is even number", number)
	} else {
		fmt.Printf("%d is odd number", number)
	}
}
