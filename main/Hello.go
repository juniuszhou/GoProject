package main

import (
	"fmt"
)

type Coin struct {
	Asset  string `json:"assets"`
	Amount int    `json:"amount"`
}

func PrintError() (result uint16) {
	fmt.Println("error")
	result = 10
	return result
}

func Call(called func() error) {
	called()
}

func QuickSort(arr []uint32) {

}

func main() {
	var a [3]int
	b := a
	a[0] = 10
	println(b[0])

	// err :=
}
