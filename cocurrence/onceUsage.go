package main

import (
	"fmt"
	"sync"
)

func Init() {
	fmt.Println("read configure here.")
}

func main() {
	// Once use to all a func just one time. later call wil be ignored.
	var once sync.Once
	once.Do(Init)
	fmt.Println("start to run after init.")
	once.Do(Init)
}
