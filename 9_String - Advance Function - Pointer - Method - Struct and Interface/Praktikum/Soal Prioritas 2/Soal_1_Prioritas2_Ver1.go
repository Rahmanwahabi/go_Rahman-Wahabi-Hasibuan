package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	// Mengkonversi input string menjadi rune slice untuk memudahkan manipulasi
	inputRunes := []rune(input)

	// Ulangi setiap rune dan terapkan offset
	for i, r := range inputRunes {
		if r == ' ' {
			continue // Skip spaces
		}

		// Geser rune berdasarkan offset, dengan mempertimbangkan kode ASCII untuk 'a' dan 'z'
		shifted := 'a' + (r-'a'+rune(offset))%('z'-'a'+1)

		// Ganti rune asli dengan yang telah digeser
		inputRunes[i] = shifted
	}

	// Mengkonversi potongan rune kembali menjadi string dan mengembalikannya
	return string(inputRunes)
}

func main() {
	// Test Case 1
	fmt.Println(caesar(3, "abc")) // def
	// Test Case 2
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	// Test Case 3
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
