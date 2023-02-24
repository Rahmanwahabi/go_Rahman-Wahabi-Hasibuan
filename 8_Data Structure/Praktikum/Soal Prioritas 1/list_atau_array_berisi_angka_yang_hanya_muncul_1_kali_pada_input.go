package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	countMap := make(map[int]int)
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		countMap[num]++
	}

	var uniqueNums []int
	for num, count := range countMap {
		if count == 1 {
			uniqueNums = append(uniqueNums, num)
		}
	}
	return uniqueNums
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
