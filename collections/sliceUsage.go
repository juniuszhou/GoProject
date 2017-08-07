package main

import "fmt"

func main() {
	// make slice and give initial size.
	var slice []int = make([]int, 10)
	fmt.Println(len(slice))

	var sub = slice[0:3]
	fmt.Println(len(sub))

}