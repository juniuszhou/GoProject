package main

func closeChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// if jobs closed and no more message then more got false.
			_, more := <-jobs
			if more {
				println("get job from main")
			} else {
				println("no more jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	// close channel
	close(jobs)
	println("over")
	<-done
}

func main() {
	closeChannel()
}
