package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i = rand.Int()
	if i > 1 {
		fmt.Println("i is big")
	}

	switch i {
	case 1:
		fmt.Println("i value is one")
		break
	case 2:
		fmt.Println("i value is two")
		break
	default:
		fmt.Println("not match")
	}
}