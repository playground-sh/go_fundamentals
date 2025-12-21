package code_organization

import (
    "fmt"
    "time"
)

// AnonymousFunction An anonymous function can be defined and executed immediately
// by placing parentheses after the function body, which can also accept arguments.
func AnonymousFunction() {
    // also known as function literals or lambda functions
    func(message string) {
        fmt.Println(message)
    }("This is an anonymous function")
}

// AssigningAnonymousFunctions You can assign an anonymous function to a variable,
// allowing you to call it later like a regular function.
func AssigningAnonymousFunctions() {
    // Assign the anonymous function to the 'greet' variable
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }

    greet("Alice")
}

// AnonymousFunctionWithGoRoutine Anonymous functions are commonly used to launch
// concurrent tasks using the go keyword.
func AnonymousFunctionWithGoRoutine() {
    go func() { // Launch an anonymous function as a goroutine
        time.Sleep(100 * time.Millisecond)
        fmt.Println("And now my watch begins")
    }()

    fmt.Println("The guards will start patrolling soon...")
    time.Sleep(200 * time.Millisecond) // Wait for the goroutine to finish
}

// counter is a closure
func counter() func() int {
    count := 0 // 'count' is defined in the outer scope
    return func() int { // The anonymous function forms a closure
        count++
        return count
    }
}

// AnonymousFunctionAsClosure A closure can access variables defined outside its own
// function body. In this example, the inner anonymous function accesses and increments
// the count variable from the outer counter function's scope.
func AnonymousFunctionAsClosure() {
    fmt.Println("An Anonymous Function as a Closure")
    c := counter()
    fmt.Println(c())
    fmt.Println(c())
}
