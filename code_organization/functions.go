package code_organization

import (
	"fmt"
	"time"
)

// add is a function with a single return value
func add(x, y int) int {
	return x + y
}

func AddDemo() {
	fmt.Println("A function with a single return value")
	var x, y = 10, 5
	var result = add(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, result)
}

// swap is a function with multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

func SwapDemo() {
	fmt.Println("A function with multiple return values")
	var a = "Hello"
	var b = "World"

	a, b = swap(a, b)
	fmt.Println(a, b)
}

// split divides a number into parts
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func SplitDemo() {
	a, b := split(17)
	fmt.Println(a, b)
}

// sum adds an arbitrary number of integer arguments and returns their sum
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func SumDemo() {
	sum()
	sum(1, 2, 3, 4, 5)

	nums := []int{10, 20, 30, 40, 50}
	sum(nums...)
}

// AnonymousFunction An anonymous function can be defined and executed immediately
// by placing parentheses after the function body, which can also accept arguments.
func AnonymousFunction() {
	// also known as function literals or lambda functions
	func(message string) {
		fmt.Println(message)
	}("This is an anonymous function")
}

// AssigningAnnonymousFunctions You can assign an anonymous function to a variable,
// allowing you to call it later like a regular function.
func AssigningAnnonymousFunctions() {
	// Assign the anonymous function to the 'greet' variable
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}

	greet("Alice")
}

// AnnonymousFunctionWithGoRoutine Anonymous functions are commonly used to launch
// concurrent tasks using the go keyword.
func AnnonymousFunctionWithGoRoutine() {
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

// AnnonymousFunctionAsClosure A closure can access variables defined outside its own
// function body. In this example, the inner anonymous function accesses and increments
// the count variable from the outer counter function's scope.
func AnnonymousFunctionAsClosure() {
	fmt.Println("An Annonymous Function as a Closure")
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
}
