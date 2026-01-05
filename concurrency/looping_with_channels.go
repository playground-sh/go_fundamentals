package concurrency

import (
	"fmt"
	"time"
)

func producer(jobs chan<- int) {
	for i := 1; i <= 10; i++ {
		jobs <- i // Send value to the channel
		time.Sleep(time.Second)
	}
	close(jobs) // Close the channel when done sending
}

func LoopingWithChannels() {
	jobs := make(chan int)
	go producer(jobs)

	// Range over the channel to receive values until it is closed
	for job := range jobs {
		fmt.Println("Received a new job:", job)
	}
}
