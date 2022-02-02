package main

import (
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if r == nil {
			println("Noting to recover")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	// panic(errors.New("this is an error"))
	b := 0
	a := 5 / b
	print(a)
}
