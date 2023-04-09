package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	// Menangani kasus bilangan kurang dari 2
	if number < 2 {
		return false
	}

	// Menangani kasus bilangan 2 dan 3
	if number == 2 || number == 3 {
		return true
	}

	// Menangani kasus bilangan genap
	if number%2 == 0 {
		return false
	}

	// Menghitung batas loop
	limit := int(math.Sqrt(float64(number))) + 1

	// Mengulangi uji pembagian dengan bilangan ganjil
	for i := 3; i < limit; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(1500450271)) // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
	fmt.Println(primeNumber(1))          // false
	fmt.Println(primeNumber(42))         // false
}
