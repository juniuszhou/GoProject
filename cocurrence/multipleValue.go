package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func SendValue(ch chan int) {
	for index := 0; index < 10; index++ {

		up, _ := rand.Int(rand.Reader,big.NewInt(100))
		fmt.Println(up)

		if up.Cmp(big.NewInt(10)) < 0 {
			return
		}
	}

	ch <- -10
}

func main() {
	ch := make(chan int)
	go SendValue(ch)
	for {
		up := <-ch
		fmt.Println(up)
		if up < 10 {
			break
		}
	}
	fmt.Println("over")
}

// fatal error: all goroutines are asleep - deadlock!
//在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的
//但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。
//所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。
//这显然是没有结果的，所以这个程序就说“算了吧，不坚持了，我自己自杀掉，报一个错给代码作者，我被deadlock了”


