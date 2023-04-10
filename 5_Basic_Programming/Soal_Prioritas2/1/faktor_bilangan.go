package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&n)

	fmt.Println("Faktor dari", n, "adalah:")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}
