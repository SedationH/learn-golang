package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			// fmt.Printf("Hi form %d\n", i)
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Printf("%v", a)
}
