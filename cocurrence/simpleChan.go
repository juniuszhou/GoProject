package main

import (
	"math/rand"
	"fmt"
)

func Call(ch chan int) {
	up := rand.Intn(10000)

	for i := 0; i < up; i += 1 {
	}
	// send value to channel
	ch <- up
}

func main(){
	// make channel array
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++{
		// make each channel
		chs[i] = make(chan int)
		// call with channel parameter
		go Call(chs[i])
	}

	for index, ch := range chs {
		// <-ch get value from channel
		fmt.Println(index, <-ch)
	}
	fmt.Println("game over")
	//close all channel
	for e := range chs {
		close(chs[e])
	}
}
