package main

import (
	"sync"
	"fmt"
	"time"
)

func UsageMutexLock(lock sync.Mutex){
	lock.Lock()
	fmt.Println("I locked the resource.")
	time.Sleep(100)
	lock.Unlock()
}

func main(){
	var lock sync.Mutex
	for index := range make([]int, 10){
		fmt.Println(index)
		go UsageMutexLock(lock)

	}
}
