package main

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	adder := getAdder()

	println(adder(1), adder(2), adder(3))
}
