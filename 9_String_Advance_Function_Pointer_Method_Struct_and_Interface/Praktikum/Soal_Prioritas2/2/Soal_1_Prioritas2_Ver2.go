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
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
