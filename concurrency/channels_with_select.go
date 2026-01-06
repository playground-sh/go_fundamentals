package concurrency

import (
	"fmt"
	"time"
)

func ChannelsWithSelectStatement() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "hello from channel 1"
	}()

	go func() {
		ch2 <- "Hi, from channel 2"
	}()

	time.Sleep(50 * time.Microsecond)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("Nothing to see here.")
	}
}
