package basic

import "fmt"

func OperatorsAndComparisons() {
	a, b := 5, 2

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a * b =", a*b)

	fmt.Println("float32(a) / float32(b) =", float32(a)/float32(b))

	fmt.Println("a == b:", a == b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
}
