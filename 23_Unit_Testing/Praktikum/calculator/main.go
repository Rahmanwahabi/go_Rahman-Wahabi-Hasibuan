package main

import (
	"fmt"
	"log"

	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/calculator/calculator"
)

func main() {
	result := calculator.Addition(3, 5)
	fmt.Println("Addition:", result)

	result = calculator.Subtraction(10, 7)
	fmt.Println("Subtraction:", result)

	result, err := calculator.Division(15, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Division:", result)

	result = calculator.Multiplication(4, 6)
	fmt.Println("Multiplication:", result)
}
