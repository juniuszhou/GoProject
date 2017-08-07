package main

import "fmt"

func main() {
	var m = make(map[int]string)
	m[0] = "Hello"
	fmt.Println(m[0])

	m[0] = "world"
	fmt.Println(m[0])
}