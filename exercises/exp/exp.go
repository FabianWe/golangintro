package main

import "fmt"

func main() {
	epsilon := 0.1
	fmt.Printf("exp(%.2f) = %.2f\n", 0.0, Exp(0.0, epsilon))
	fmt.Printf("exp(%.2f) = %.2f\n", 1.0, Exp(1.0, epsilon))
	fmt.Printf("exp(%.2f) = %.2f\n", 4.2, Exp(4.2, epsilon))
}

func Exp(x, epsilon float64) float64 {
	// Your code here
}
