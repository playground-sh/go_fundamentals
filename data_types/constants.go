package data_types

import "fmt"

func Constants() {
	const g = 9.81
	fmt.Println("The Gravitational Acceleration on planet Earth is:", g)

	fmt.Println("Please don't change that. I like it the way it is.")
	// g = 10 // uncommenting this line will produce a compiler error
}
