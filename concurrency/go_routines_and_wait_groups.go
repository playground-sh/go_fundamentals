package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutinesSimpleDemo() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("<This is asynchronous>")
		wg.Done()
	}()

	go func() {
		fmt.Println("[This is asynchronous, too...]")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Now we are back to normal")
}

func worker(id int, wg *sync.WaitGroup) {
	// Schedule the Done call to happen when the worker function finishes
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// fake work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d completed its mission\n", id)
}

func GoRoutinesMoreDemo() {
	// A WaitGroup is used to wait for all goroutines to finish
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter for each goroutine launched
		wg.Add(1)

		// Launch a new goroutine using the 'go' keyword
		go worker(i, &wg)
	}

	// Block the main function until the WaitGroup counter is zero
	wg.Wait()
	fmt.Println("Our mission here is complete, all workers have returned")
}
