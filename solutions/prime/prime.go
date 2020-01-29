package main

import (
	"fmt"
)

func main() {
	numPrimes := 0
	var i uint = 2
	for {
		if IsPrime(i) {
			numPrimes++
			fmt.Println(i)
		}
		if numPrimes == 100 {
			break
		}
		i++
	}
}

// IsPrime tests if n is a prime number.
func IsPrime(n uint) bool {
	var i uint = 2
	for ; i * i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
