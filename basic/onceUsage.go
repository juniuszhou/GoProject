package main

import (
	"fmt"
	"sync"
)

func Init(){
	fmt.Println("read configure here.")
}

func main(){
	var once sync.Once
	once.Do(Init)
	fmt.Println("start to run after init.")
}

