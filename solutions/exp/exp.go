package main

import "fmt"

func main() {
	epsilon := 0.1
	fmt.Printf("exp(%.2f) = %.2f\n", 0.0, Exp(0.0, epsilon))
	fmt.Printf("exp(%.2f) = %.2f\n", 1.0, Exp(1.0, epsilon))
	fmt.Printf("exp(%.2f) = %.2f\n", 4.2, Exp(4.2, epsilon))
}

func Exp(x, epsilon float64) float64 {
	num := 1.0
	den := 1.0
	sum := 1.0
	i := 1.0
	for {
		num *= x
		den *= i
		i += 1.0
		newSum := sum + num/den
		diff := newSum - sum
		if diff < 0 {
			diff *= -1.0
		}
		if diff < epsilon {
			return newSum
		}
		sum = newSum
	}
}
