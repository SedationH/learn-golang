package main

import "learn-golang/function/fib"

func main() {
	fib := fib.GetFibFn()
	println(fib(), fib(), fib(), fib())
}
