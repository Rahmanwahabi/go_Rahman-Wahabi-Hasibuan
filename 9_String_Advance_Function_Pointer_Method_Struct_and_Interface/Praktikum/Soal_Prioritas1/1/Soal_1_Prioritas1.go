package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // bahan bakar dalam liter
}

func (c Car) EstimateRange() float64 { // konversi bahan bakar dari liter ke milliliter

	fuel := c.FuelIn * 1000.0 // hitung jarak yang bisa ditempuh berdasarkan bahan bakar yang tersedia

	rangeInKM := fuel * 1.5
	return rangeInKM
}

func main() {
	c := Car{
		Type:   "Tesla",
		FuelIn: 45.0, // 45 liter bahan bakar
	}
	rangeInKM := c.EstimateRange()
	fmt.Printf("Perkiraan jarak yang bisa ditempuh: %.2f KM\n", rangeInKM)
}
