package main

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for message := range queue {
		println(message)
	}
}
