package main

import "fmt"

func generate(numRows int) [][]int {
	// Membuat array kosong untuk menyimpan segitiga Pascal
	pascalTriangle := make([][]int, numRows)

	// Switch case untuk menangani kasus numRows yang berbeda
	switch numRows {
	case 0:
		return pascalTriangle
	case 1:
		pascalTriangle[0] = []int{1}
		return pascalTriangle
	default:
		pascalTriangle[0] = []int{1}
		for i := 1; i < numRows; i++ {
			newRow := make([]int, i+1)
			newRow[0], newRow[i] = 1, 1
			for j := 1; j < i; j++ {
				newRow[j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
			}
			pascalTriangle[i] = newRow
		}
		return pascalTriangle
	}
}

func main() {
	pascalTriangle := generate(5)
	fmt.Println(pascalTriangle)
	// Output: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
