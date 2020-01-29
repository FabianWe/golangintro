package main

import (
	"fmt"
	// "strings" // You probably want this...
)

func main() {
	for _, s := range []string{"foo@oof", "Regalnebenlager", "tattarrattat", "a", "foo", "abc"} {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", s, IsPalindrome(s))
	}
}

func IsPalindrome(s string) bool {
	// Your code here
}
