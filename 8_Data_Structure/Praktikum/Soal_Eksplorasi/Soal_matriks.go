package main

import (
	"fmt"
	"math"
)

func main() {
	// Input matriks persegi
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	// Menghitung selisih absolut antara jumlah diagonal
	diff := int(math.Abs(float64(matrix[0][0] + matrix[1][1] + matrix[2][2] - matrix[0][2] - matrix[1][1] - matrix[2][0])))

	// Menampilkan hasil
	fmt.Println("Selisih absolut antara jumlah diagonal: ", diff)
}
