package composite

import "fmt"

func SetsDemo() {
	pizzas := NewSet[string]()
	pizzas.Add("Pepperoni")
	pizzas.Add("Cheese")
	pizzas.Add("Cheese")	// adding duplicate
	pizzas.Add("Cheese")	// adding duplicate
	pizzas.Add("Cheese")	// adding duplicate
	pizzas.Add("Classic")
	fmt.Println("Set of Pizzas (String):", pizzas)

	roads := NewSet[int]()
	roads.Add(1)
	roads.Add(3)
	roads.Add(5)
	roads.Add(7)
	roads.Add(1) // adding duplicate
	roads.Add(3) // adding duplicate
	fmt.Println("Set of Numbers (int):",roads)
}
