package main

import "fmt"

func Loop() {
	var values []int
	for value := range values {
		fmt.Println(value)
	}

	// use for as while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// use for as for
	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}

	// use for as loop until
	for {
		fmt.Println(i)
		if i > 4 {
			break
		}
	}

	for i := 0; i < 10; i += 1 {
		fmt.Print(i + ' ')
	}

	var a [10]int
	for index, value := range a {
		fmt.Println(index, value)
	}

}
