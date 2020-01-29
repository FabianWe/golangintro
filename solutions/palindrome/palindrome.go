package main

import (
	"fmt"
	"strings"
)

func main() {
	for _, s := range []string{"foo@oof", "Regalnebenlager", "tattarrattat", "a", "foo", "abc"} {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", s, IsPalindrome(s))
	}
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	slice := []rune(s)
	forward, backward := 0, len(s) - 1
	for forward < backward {
		if slice[forward] != slice[backward] {
			return false
		}
		forward++
		backward--
	}
	return true
}
