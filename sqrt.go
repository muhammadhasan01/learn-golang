package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	l, r := float64(0), x
	for l < r {
		mid := (l + r) / 2.0
		if mid*mid <= x {
			l = mid + 0.0000001
		} else {
			r = mid - 0.0000001
		}
	}
	return l
}

func main() {
	fmt.Println(sqrt(2), math.Sqrt(2))
}
