package main

import "fmt"

func main() {
	var nilai float64
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	var grade string
	if nilai < 0 || nilai > 100 {
		grade = "Nilai Invalid"
	} else if nilai >= 80 {
		grade = "A"
	} else if nilai >= 65 {
		grade = "B"
	} else if nilai >= 50 {
		grade = "C"
	} else if nilai >= 35 {
		grade = "D"
	} else if nilai >= 0 {
		grade = "E"
	}

	fmt.Println("Grade Anda adalah:", grade)
}
