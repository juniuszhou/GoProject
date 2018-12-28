package main

import (
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops int64 = 0
	for i := 0; i < 30; i++ {
		go func() {
			for {
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	result := atomic.LoadInt64(&ops)
	println(result)
}
