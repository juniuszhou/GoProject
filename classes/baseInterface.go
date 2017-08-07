package main

import "fmt"

type BaseInterface interface {
	Base()
}

type BaseImplement struct {

}

func (base *BaseImplement) Base() {
	fmt.Println("implemented Base method.")
}

func main() {
	var b BaseInterface = new(BaseImplement)
	b.Base()
}
