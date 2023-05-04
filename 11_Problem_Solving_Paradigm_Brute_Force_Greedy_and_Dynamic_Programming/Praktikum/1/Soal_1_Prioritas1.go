package main

import (
	"fmt"
	"strconv"
)

func binaryArray(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		binaryStr := strconv.FormatInt(int64(i), 2)
		val, _ := strconv.ParseInt(binaryStr, 10, 64)
		ans[i] = int(val)
	}
	return ans
}

func main() {
	fmt.Println(binaryArray(2)) // [0 1 10]
	fmt.Println(binaryArray(3)) // [0 1 10 11]
}
