package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// use argument here, common mistake!
		go func(val int) {
			fmt.Println(val)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
