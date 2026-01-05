package error_handling

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileOps(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		message := fmt.Sprintf("fatal: failed to open file: %v", err)
		fmt.Println(message)
	}

	return file
}

func CreatePanic() {
	// panic!
	file := fileOps("does_not_exist.txt")

	// nothing below will execute
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			message := fmt.Errorf("failed to close file: %v", err)
			fmt.Println(message)
		}
	}(file)

	buffer := make([]byte, 100)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}
}
