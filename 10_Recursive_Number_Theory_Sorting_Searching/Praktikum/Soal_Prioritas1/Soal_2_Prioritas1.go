package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	countMap := make(map[string]int)

	// hitung kemunculan setiap item
	for _, item := range items {
		countMap[item]++
	}

	// buat slice dari pair
	pairs := make([]pair, 0, len(countMap))
	for name, count := range countMap {
		pairs = append(pairs, pair{name, count})
	}

	// sorting slice berdasarkan jumlah kemunculan dari terkecil ke terbesar
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// output: [{golang 1} {ruby 2} {js 4}]

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// output: [{C 1} {D 2} {B 3} {A 4}]

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// output: [{football 1} {basketball 1} {tenis 1}]
}
