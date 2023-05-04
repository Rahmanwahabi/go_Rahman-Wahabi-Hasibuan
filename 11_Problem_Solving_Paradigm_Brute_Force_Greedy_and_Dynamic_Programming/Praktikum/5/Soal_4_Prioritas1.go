package main

import (
	"fmt"
	"math"
)

// disini saya pake devide and conquer , kalo brute force terlalu lama
func findXYZ(A, B, C int, min, max int) (int, int, int) {
	if max < min {
		return 0, 0, 0
	}

	if max == min {
		if A == 3*max && B == max*max*max {
			return max, max, max
		}
		return 0, 0, 0
	}

	mid := (max + min) / 2

	x1, y1, z1 := findXYZ(A, B, C, min, mid)
	if x1 != 0 && y1 != 0 && z1 != 0 {
		return x1, y1, z1
	}

	x2, y2, z2 := findXYZ(A, B, C, mid+1, max)
	if x2 != 0 && y2 != 0 && z2 != 0 {
		return x2, y2, z2
	}

	for i := min; i <= mid; i++ {
		for j := mid + 1; j <= max; j++ {
			k := A - i - j
			if i == j || j == k || k == i {
				continue
			}
			if i*j*k == B && i*i+j*j+k*k == C {
				return i, j, k
			}
		}
	}

	return 0, 0, 0
}

func SimpleEquation(A, B, C int) {
	x, y, z := findXYZ(A, B, C, 1, int(math.Sqrt(float64(C))))
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("no solution")
	} else {
		fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
	}
}

func main() {
	SimpleEquation(1, 2, 3)     // no solution
	SimpleEquation(6, 6, 14)    // 1 2 3
	SimpleEquation(10000, 1, 1) // no solution
}
