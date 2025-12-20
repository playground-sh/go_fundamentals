package code_organization

import "fmt"

// add is a function with a single return value
func add(x, y int) int {
	return x + y
}

func AddDemo() {
	var x, y = 10, 5
	var result = add(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, result)
}

// swap is a function with multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

func SwapDemo() {
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

func AnonymousFunction() {
	// also known as function literals or lambda functions
	func(message string) {
		fmt.Println(message)
	}("This is an anonymous function")
}
