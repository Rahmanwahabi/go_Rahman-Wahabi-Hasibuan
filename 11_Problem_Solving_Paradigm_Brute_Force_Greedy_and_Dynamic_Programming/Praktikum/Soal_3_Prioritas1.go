package main

import "fmt"

func FBI(n int, memo map[int]int) int {
	// Cek apakah n sudah pernah dihitung sebelumnya
	if val, ok := memo[n]; ok {
		return val
	}
	// Base case
	if n == 0 || n == 1 {
		return n
	}
	// Hitung nilai Fibonacci ke-n dengan rekursi
	memo[n] = FBI(n-1, memo) + FBI(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	fmt.Println(FBI(0, memo))  // Output: 0
	fmt.Println(FBI(1, memo))  // Output: 1
	fmt.Println(FBI(2, memo))  // Output: 1
	fmt.Println(FBI(3, memo))  // Output: 2
	fmt.Println(FBI(5, memo))  // Output: 5
	fmt.Println(FBI(6, memo))  // Output: 8
	fmt.Println(FBI(7, memo))  // Output: 13
	fmt.Println(FBI(9, memo))  // Output: 34
	fmt.Println(FBI(10, memo)) // Output: 55

}
