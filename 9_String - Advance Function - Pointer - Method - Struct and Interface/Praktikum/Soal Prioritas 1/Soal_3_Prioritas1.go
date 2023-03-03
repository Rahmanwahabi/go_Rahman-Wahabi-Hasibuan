package main

import (
	"fmt"
)

func Compare(a, b string) string {
	var result string
	lenA := len(a)
	lenB := len(b)
	minLen := lenA
	if lenB < lenA {
		minLen = lenB
	}
	for i := 0; i < minLen; i++ {
		switch {
		case a[i:i+1] == b[i:i+1]:
			result += a[i : i+1]
		default:
			break
		}
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
