package main

import "fmt"

func getFunc() {
	f := func(x int, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))
}

func closure() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	getFunc()
	fmt.Println("++++++++++++++++++++")
	c1 := closure()
	fmt.Println(c1(), c1(), c1())
	c2 := closure()
	fmt.Println(c2(), c2())
}
