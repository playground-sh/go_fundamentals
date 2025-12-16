package control_flow

import (
	"fmt"
	"os"
)

func DeferredFileIO() {
	file, err := os.Open("C:\\main.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Defer closing the file. This ensures file.Close() is called
	// when the function returns, regardless of errors.
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Red %d bytes:\n%s\n", n, string(buffer[:n]))
}

func StackingDeferredCalls() {
	// multiple defer calls are called in the LIFO order, like a stack
	fmt.Println("Start stack")
	defer fmt.Println("First Defer")
	defer fmt.Println("Second Defer")
	defer fmt.Println("Third Defer")
	fmt.Println("End stack")
	fmt.Println("Deferred calls to `Println()` will appear in LIFO order")
}
