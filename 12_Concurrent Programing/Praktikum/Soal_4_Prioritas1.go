package main

import (
	"fmt"
	"sort"
	"sync"
)

func calculateFactorial(n int, wg *sync.WaitGroup, mu *sync.Mutex, results map[int]int) {
	defer wg.Done()

	factorial := 1

	for i := 1; i <= n; i++ {
		mu.Lock()
		factorial *= i
		mu.Unlock()
	}

	mu.Lock()
	results[n] = factorial
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	numbers := []int{1, 4, 3, 2, 5}
	results := make(map[int]int)

	for _, n := range numbers {
		wg.Add(1)
		go calculateFactorial(n, &wg, &mu, results)
	}

	wg.Wait()

	// Mengurutkan map berdasarkan key
	keys := make([]int, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Cetak hasil faktorial
	for _, k := range keys {
		fmt.Printf("Faktorial dari %d adalah %d\n", k, results[k])
	}

	fmt.Println("Selesai menghitung faktorial.")
}
