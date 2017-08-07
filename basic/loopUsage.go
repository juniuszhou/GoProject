package main

import "fmt"

// array is fixed size collection. length assigned at init
func main(){
	for i := 0; i < 10; i += 1 {
		fmt.Print(i + ' ')
	}

	var a [10]int
	for index, value := range a{
		fmt.Println(index, value)
	}



}
