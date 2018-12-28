package main

import (
	"fmt"
	"math/rand"
)

func Loop(index int) {
	up := rand.Int()
	for i := 0; i < up; i += 1 {

	}
	fmt.Println(index, up)
}

func main() {
	for i := 0; i < 10; i += 1 {
		go Loop(i)
	}
}
