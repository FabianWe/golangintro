package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(MultiplySlice(s, 3))
}

func MultiplySlice(s []int, num int) []int {
	res := make([]int, 0, num*len(s))
	for i := 0; i < num; i++ {
		// there is also a shorter version:
		// res = append(res, s...)
		for _, val := range s {
			res = append(res, val)
		}
	}
	return res
}
