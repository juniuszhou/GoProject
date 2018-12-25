package main

import "fmt"

type ICall interface {
	CallFunc()
}

type Implement struct {
}

func (i *Implement) CallFunc() {
	fmt.Println("Implementation called.")
}
