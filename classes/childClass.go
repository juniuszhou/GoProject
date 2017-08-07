package main

import "fmt"
import "github.com/juniuszhou/GoProject/classes/"

type Child struct {
	Base
	ChildName string
}

func main() {
	child := new(Child)
	fmt.Println(child.ChildName)
}