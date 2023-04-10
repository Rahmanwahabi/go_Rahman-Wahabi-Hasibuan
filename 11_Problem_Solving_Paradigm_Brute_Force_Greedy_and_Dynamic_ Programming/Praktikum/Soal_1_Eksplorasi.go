package main

import "fmt"

func Romanum(num int) string {
	romanMap := []struct {
		roman string
		value int
	}{
		{"M", 1000}, {"CM", 900}, {"D", 500}, {"CD", 400}, {"C", 100}, {"XC", 90}, {"L", 50},
		{"XL", 40}, {"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4}, {"I", 1},
	}

	roman := ""
	for _, rm := range romanMap {
		for num >= rm.value {
			roman += rm.roman
			num -= rm.value
		}
	}
	return roman
}

func main() {
	fmt.Println(Romanum(4))    // Output: IV
	fmt.Println(Romanum(9))    // Output: IX
	fmt.Println(Romanum(23))   // Output: XXIII
	fmt.Println(Romanum(2021)) // Output: MMXXI
	fmt.Println(Romanum(1646)) // Output: MDCXLVI
}
