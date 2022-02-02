package main

import (
	"bufio"
	"fmt"
	"learn-golang/function/fib"
	"os"
)

func tryDefer() {
	println(1)
	defer println(2)
	defer println(3)

	for i := 0; i < 100; i++ {
		defer println(i)
		if i == 30 {
			// panic("too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		// TODO:对这里的处理机制不清晰
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s %s %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fib := fib.GetFibFn()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	writeFile("fib.txt")

	// tryDefer()
}
