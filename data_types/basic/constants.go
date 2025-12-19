package basic

import "fmt"

func Constants() {
	const g = 9.81
	fmt.Println("The Gravitational Acceleration on planet Earth is:", g)

	fmt.Println("Please don't change that. I like it the way it is.")
	// g = 10 // uncommenting this line will produce a compiler error
}

func Iota() {
	const x = iota
	fmt.Println("x = iota:", x)

	// `iota` resets in a new scope
	const y = iota
	fmt.Println("y = iota:", y)

	// but is incremented in a block scope
	const (
		a = iota
		b
		c
	)
	fmt.Println("a:", a, ", b:", b, ", c:", c)

	// and the value of iota depends on its position
	const (
		d = 5 * 20
		e
		f = iota
		g
		h = iota * 5
	)
	fmt.Println("d:", d, ", e:", e, ", f:", f, ", g:", g, ", h:", h)
}
