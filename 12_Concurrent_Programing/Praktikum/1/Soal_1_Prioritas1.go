package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}
}

func main() {
	go func() {
		for {
			printMultiples(7)
			time.Sleep(3 * time.Second)
		}
	}()
	var input string
	fmt.Scanln(&input)
}

//untuk menghentikan loopnya cukup tekan ctrl + c
