package main

import "fmt"

func main() {
	var up, down, high float64

	fmt.Print("Enter the length of the top of the trapezoid: ")
	fmt.Scan(&up)

	fmt.Print("Enter the length of the down side of the trapezoid: ")
	fmt.Scan(&down)

	fmt.Print("Enter the trapezoid height: ")
	fmt.Scan(&high)

	area := (up + down) * high / 2

	fmt.Printf("The area of the trapezoid is: %.2f\n", area)
}
