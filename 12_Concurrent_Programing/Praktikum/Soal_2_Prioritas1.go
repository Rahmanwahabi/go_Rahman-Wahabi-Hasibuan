package main

import (
	"fmt"
)

func printMultiples(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 3 // Mengirim bilangan kelipatan 3 ke channel
	}
	close(ch) // Menutup channel setelah selesai mengirim bilangan
}

func main() {
	ch := make(chan int)  // Membuat channel dengan tipe int
	go printMultiples(ch) // Menjalankan goroutine untuk mencetak bilangan kelipatan 3

	// Mengambil nilai dari channel dan mencetaknya
	for multiple := range ch {
		fmt.Println(multiple)
	}
}
