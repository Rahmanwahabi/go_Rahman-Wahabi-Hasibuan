package main

import "fmt"

func main() {
	var input string
	fmt.Print("input kata: ")
	fmt.Scanln(&input)

	if Palindrom(input) {
		fmt.Printf("%s adalah palindrom.\n", input)
	} else {
		fmt.Printf("%s adalah tak palindrom.\n", input)
	}
}

func Palindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
