package main

import "fmt"

func main() {
	var normalChan chan int
	var writeChan chan<- int
	var readChan <-chan int

	var result = <-normalChan
	fmt.Println(result)
	writeChan <- 10 // only for write
	_ = <-readChan  // only for read

	// transform
	ch := make(chan int)
	_ = <-ch
	readableChan := (<-chan int)(ch)
	_ = <-readableChan

	// check after close
	close(ch)
	value, ok := <-ch
	if ok {
		fmt.Println(value)
	}

}
