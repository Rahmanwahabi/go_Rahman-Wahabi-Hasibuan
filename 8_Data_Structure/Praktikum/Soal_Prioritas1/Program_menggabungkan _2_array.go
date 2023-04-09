package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	//  Buat map untuk menyimpan elemen-elemen pada arrayA
	uniqueElements := make(map[string]bool)

	//  Iterasi elemen-elemen pada arrayA dan tambahkan ke map
	for _, a := range arrayA {
		uniqueElements[a] = true
	}

	// Iterasi elemen-elemen pada arrayB dan tambahkan ke map jika belum ada
	for _, b := range arrayB {
		if !uniqueElements[b] {
			uniqueElements[b] = true
		}
	}

	// Buat slice untuk menyimpan elemen-elemen unik
	mergedArray := make([]string, 0, len(uniqueElements))

	// Iterasi key pada map dan tambahkan ke slice, dimulai dari arrayA terlebih dahulu
	for key := range uniqueElements {
		mergedArray = append(mergedArray, key)
	}

	return mergedArray
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))          // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))                         // ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})) // ["alisa", "yoshimitsu","devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))                                          // ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))                                                     // ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))                                                               // []
}
