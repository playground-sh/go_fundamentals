package error_handling

import (
    "errors"
    "fmt"
    "log"
    "os"
)

func CreatingErrors() {
    err := errors.New("well, I didn't see this coming")
    fmt.Println(err)

    wrappedError := fmt.Errorf("this error wraps around the first error: %w", err)
    fmt.Println(wrappedError)
}

func ErrorHandling() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    // the following code will not execute
    defer func(file *os.File) {
        err := file.Close()
        if err != nil {
            fmt.Println("Failed to close file")
        }
    }(file)

    fmt.Println("File opened successfully!")
}
