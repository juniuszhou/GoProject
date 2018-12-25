package main

import "fmt"

type Base struct {
	BaseName string
}

type Child struct {
	Base
	ChildName string
}

func main() {
	child := new(Child)
	child.ChildName = "hello"
	child.BaseName = "world"
	fmt.Println(child.ChildName)
}
