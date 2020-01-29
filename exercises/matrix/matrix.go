package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// initialize random engine
	rand.Seed(time.Now().UTC().UnixNano())
	matricesA := make([]Matrix, 10)
	matricesB := make([]Matrix, 10)
	n := 500
	m := 500
	k := 600
	for i := 0; i < 10; i++ {
		matricesA[i] = NewRandomMatrix(n, m)
		matricesB[i] = NewRandomMatrix(m, k)
	}
	start := time.Now()
	for _, m1 := range matricesA {
		for _, m2 := range matricesB {
			m1.MultSeq(m2)
		}
	}
	end := time.Now()
	fmt.Println("Sequential time", end.Sub(start))
	start = time.Now()
	for _, m1 := range matricesA {
		for _, m2 := range matricesB {
			m1.MultConc(m2)
		}
	}
	end = time.Now()
	fmt.Println("Concurrent time", end.Sub(start))
}

type Matrix [][]float64

func NewMatrix(n, m int) Matrix {
	var res [][]float64
	for i := 0; i < n; i++ {
		res = append(res, make([]float64, m))
	}
	return res
}

func NewRandomMatrix(n, m int) Matrix {
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		col := make([]float64, m)
		for j := 0; j < m; j++ {
			col[j] = float64(rand.Intn(50))
		}
		res[i] = col
	}
	return res
}

func (a Matrix) Compare(b Matrix) bool {
	delta := 0.0001
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if math.Abs(a[i][j]-b[i][j]) > delta {
				return false
			}
		}
	}
	return true
}

func (a Matrix) MultSeq(b Matrix) Matrix {
	// assume correct dimensions and well-formed matrix
	// also no rows is not allowed
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	res := NewMatrix(n, k)
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			sum := 0.0
			for pos := 0; pos < m; pos++ {
				sum += a[i][pos] * b[pos][j]
			}
			res[i][j] = sum
		}
	}
	return res
}

func (a Matrix) MultConc(b Matrix) Matrix {
	// Your code here
}

