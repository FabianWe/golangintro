package tests

import (
	"golangsrc/examples/factorial"
	"testing"
)

func TestFactorial(t *testing.T) {
	values := []uint{0, 1, 2, 3}
	results := []uint{1, 1, 2, 6}
	// change to this results to produce an error
	// results := []uint{1, 1, 2, 7}

	for i := 0; i < len(values); i++ {
		f := factorial.Factorial(values[i])
		if f != results[i] {
			t.Errorf("Expected %d! = %d, got %d", values[i], results[i], f)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.Factorial(20)
	}
}

func BenchmarkFactorialRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.FactorialRec(20)
	}
}
