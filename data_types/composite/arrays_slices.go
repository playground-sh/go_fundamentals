package composite

import (
	"fmt"
	"strings"
)

func ArraysAndSlices() {
	var oddNumbers = [5]int{1, 3, 5, 7, 9}
	fmt.Println("Odd Numbers", oddNumbers)

	slicesOfPizza := []string{"Classic"}
	slicesOfPizza = append(slicesOfPizza, "Cheese Pizza", "Pepperoni Pizza")
	pizzas := strings.Join(slicesOfPizza, ", ")
	fmt.Println("[" + pizzas + "]")
}
