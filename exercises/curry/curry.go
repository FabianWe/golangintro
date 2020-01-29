package main

import "fmt"

func main() {
	adder := func(x, y int) int {
		return x + y
	}
	fmt.Printf("Curry(AddNumber, 15)(27) = %d\n", Curry(adder, 15)(27))
}

func Curry(f func(x, y int) int, x int) func(y int) int {
	return func(y int) int {
		return f(x, y)
	}
}
