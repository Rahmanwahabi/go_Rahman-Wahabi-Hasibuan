package main

import (
	"fmt"
	"strings"
	"sync"
)

func countFrequency(text string, wg *sync.WaitGroup, result chan map[rune]int) {
	defer wg.Done()

	// Inisialisasi map yang menyimpan frekuensi huruf
	frequencies := make(map[rune]int)

	// Menghitung frekuensi huruf dalam teks
	for _, char := range text {
		frequencies[char]++
	}

	result <- frequencies
}

func main() {
	// Teks yang akan dihitung frekuensi hurufnya
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	// Membuat slice dari teks untuk diolah secara paralel
	textSlice := strings.Split(text, "")

	// Inisialisasi channel untuk menampung hasil perhitungan
	result := make(chan map[rune]int)

	// Inisialisasi WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup
	wg.Add(len(textSlice))

	// Memproses setiap bagian teks secara paralel
	for _, char := range textSlice {
		wg.Add(1)
		go countFrequency(char, &wg, result)
	}

	// Mengumpulkan hasil perhitungan dari setiap goroutine
	frequencies := make(map[rune]int)
	go func() {
		for freq := range result {
			for k, v := range freq {
				frequencies[k] += v
			}
			wg.Done()
		}
	}()

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Mencetak hasil perhitungan
	for k, v := range frequencies {
		fmt.Printf("%c : %d\n", k, v)
	}
}
