package main

import "fmt"

type MyError struct {
	Err error
}

func (e *MyError) Error() string {
	fmt.Println("Error called.")
	return "MyError"
}

func main(){
	var error = new(MyError)
	error.Error()

}