package main

import "fmt"

// panic 是用来表示非常严重的不可恢复的错误的
func main(){

	defer func(){
		fmt.Println("defer called.")
		//recover panic. if not caught in recover, then panic will pass to whole call stack.
		//the interface instance thrown in panic will be caught via recover
		r := recover()
		if r != nil {
			fmt.Println("run time error caught.", r)
		}
	}()

	panic(2)

	defer func(){
		fmt.Println("defer after panic")
	}()


}