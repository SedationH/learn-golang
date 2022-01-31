package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// convertToBin
// 2 => 10
// 8 => 1000
func convertToBin(n int) string {
	binStr := ""
	for ; n > 0; n /= 2 {
		binStr = strconv.Itoa(n%2) + binStr
	}
	return binStr
}

func printFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	println(
		convertToBin(2),
		convertToBin(8),
	)

	printFile("/Users/sedationh/workspace/study/learn-golang/basic/branch/c.txt")
}
