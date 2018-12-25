package main

func getFunc() {
	f := func(x int, y int) int {
		return x + y
	}
	println(f(1, 2))
}

func main() {
	getFunc()
}
