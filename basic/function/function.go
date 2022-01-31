package main

func useFunc(fn func(int, int) (int, int), a, b int) {
	fn(a, b)
}

func main() {
	useFunc(func(a, b int) (int, int) {
		println("use useFunc", a, b)
		return 1, 2
	}, 1, 2)

}
