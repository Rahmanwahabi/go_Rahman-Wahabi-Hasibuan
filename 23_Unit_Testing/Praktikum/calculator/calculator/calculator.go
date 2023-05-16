package calculator

import "errors"

// Addition menjumlahkan dua bilangan
func Addition(a, b float64) float64 {
	return a + b
}

// Subtraction mengurangkan bilangan kedua dari bilangan pertama
func Subtraction(a, b float64) float64 {
	return a - b
}

// Division membagi bilangan pertama dengan bilangan kedua
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Multiplication mengalikan dua bilangan
func Multiplication(a, b float64) float64 {
	return a * b
}
