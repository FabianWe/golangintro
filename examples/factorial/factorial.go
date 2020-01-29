package factorial

func Factorial(n uint) uint {
	var res uint = 1
	var i uint = 1
	for ; i <= n; i++ {
		res *= i
	}
	return res
}

func FactorialRec(n uint) uint {
	if n == 0 {
		return 1
	} else {
		return n * FactorialRec(n-1)
	}
}
