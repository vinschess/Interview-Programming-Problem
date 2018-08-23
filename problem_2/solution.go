package main

import (
	"fmt"
	"math"
)

func sqrt(val float64) float64 {
	sqrt_val := val / 2
	prev := 0.0
	for sqrt_val-prev >= 0.001 || prev-sqrt_val >= 0.001 {
		prev = sqrt_val
		sqrt_val -= (sqrt_val*sqrt_val - val) / (2 * sqrt_val)
	}
	return math.Round(sqrt_val*100) / 100
}

func main() {
	fmt.Println(sqrt(1))
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(25))
	fmt.Println(sqrt(100))
	fmt.Println(sqrt(101))
	fmt.Println(sqrt(176))
}
