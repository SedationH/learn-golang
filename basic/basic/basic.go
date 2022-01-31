package main

import (
	"fmt"
	"math"
)

func zeroValue() {
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // ""

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var a interface{} // nil

	fmt.Printf("%v", a)
}

func typeDeduction() {

	var a, b = 1.0, 2
	c := 1.0
	fmt.Printf("%T %T %v", a, b, c) // float64 int
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func triangle() {
	a, b := 3, 4
	c := calcTriangle(a, b)
	println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, bb    = 3, 4
	)
	// 注意不需要转 float64 了，理解为文本替换
	c := int(math.Sqrt(a*a + bb*bb))
	println(c)

	const (
		cpp = iota
		_
		python
		golang
	)
	println(cpp, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	println(
		bb,
		kb,
		mb,
		gb,
		tb,
		pb,
	)
}

func main() {
	fmt.Println("Hello SedationH")

	/*
		String and slice of bytes (treated equivalently with these verbs):
				%s	the uninterpreted bytes of the string or slice
				%q	a double-quoted string safely escaped with Go syntax
				%x	base 16, lower-case, two characters per byte
				%X	base 16, upper-case, two characters per byte
	*/
	zeroValue()
	typeDeduction()
	triangle()
	consts()
}
