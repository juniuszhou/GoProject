package main

import "time"

func blockChannel() {
	message := make(chan string)

	select {
	case msg := <-message:
		println("got message " + msg)
	case <-time.After(time.Second * 2):
		println("timeout after 2 seconds.")
	}
}

func nonBlockChannel() {
	message := make(chan string)

	select {
	case msg := <-message:
		println("got message " + msg)
	case <-time.After(time.Second * 2):
		println("timeout after 2 seconds.")
	default:
		println("no message, match the default case and exit select")
	}
}
func main() {
	blockChannel()
	nonBlockChannel()

}
