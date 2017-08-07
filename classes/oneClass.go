package main

import "fmt"

type  oneClass struct {
	name string

}


func (one oneClass) GetName()string {
	return one.name
}

// can declare it since it is duplicate with call get name via object
//func (one * oneClass) GetName() string {
//	return one.name
//}

func GetName(one * oneClass)string {
	return one.name
}

func GetObjectName(one oneClass)string {
	return one.name
}

func main(){
	var one = new(oneClass)
	fmt.Println(one.GetName())

	//call via pointer
	fmt.Println(GetName(one))
	//call via object
	fmt.Println(*one)


	var two = *(new(oneClass))
	fmt.Println(GetObjectName(two))
	fmt.Println(GetName(&two))

}