package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a[0])

	for index, value := range a {
		fmt.Println(index, value)
	}

	a2 := [...]int{7, 77, 777, 7777, 77777}
	for index := 0; index < len(a2); index++ {
		fmt.Println(a2[index])
	}

	// pointer array
	pa := [5]*int{0: new(int), 1: new(int)}
	fmt.Println(pa[0]) // print address of first item
	fmt.Println(*pa[0])

}
