package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// buat sebuah map untuk menyimpan hitungan setiap string
	counts := make(map[string]int)

	// Lakukan iterasi pada slice dan hitung setiap string
	for _, str := range slice {
		counts[str]++
	}

	return counts
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
