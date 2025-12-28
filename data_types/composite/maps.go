package composite

import "fmt"

func MapsDemo() {
	// maps are key value pairs
	var emptyMap map[string][]int
	fmt.Println("emptyMap:", emptyMap)
	fmt.Println()

	var beverages map[string][]string
	beverages = map[string][]string{
		"Coffee": {"Espresso", "Cappuccino", "Latte"},
		"Tea":    {"Ginger Tea", "Lemon Crush", "Malai Cha"},
	}
	fmt.Println(beverages)
	fmt.Println(beverages["Tea"]) // access members by key
	fmt.Println(beverages["Coffee"])
	fmt.Println()

	// add new key-value pairs
	beverages["Drinks"] = []string{"Cola", "Lime Soda", "Melon Juice"}
	fmt.Println("beverages after addition: ")
	fmt.Println(beverages)
	fmt.Println()

	// remove an existing pair
	delete(beverages, "Coffee")
	fmt.Println("beverages after removal: ")
	fmt.Println(beverages)
	fmt.Println()

	// accessing a key after removal doesn't throw error
	// printing it will print default empty value for its type of the value item
	fmt.Println(beverages["Coffee"])
	fmt.Println()

	// since it will print/return even empty values as a value, to be sure whether it was the
	// actual value set for the key or it's' empty because it was deleted, we use the following
	// pattern. if `ok` returs `false`, the value was deleted/not present
	val, ok := beverages["Coffee"]
	fmt.Println(val, ok)
	fmt.Println()

	bevvy := beverages
	bevvy["Coffee"] = []string{"Honey Latte"} // modifying one map
	fmt.Println(bevvy)
	fmt.Println()
	fmt.Println(beverages) // affects the other
	fmt.Println()
}

func MapsIterationDemo() {
	fruitsPrices := map[string]float32{
		"Banana": 120.0,
		"Orange": 180.0,
		"Apples": 220.5,
		"Grapes": 375.25,
		"Pears":  410.25,
	}

	for fruit, price := range fruitsPrices {
		fmt.Println(fruit, "costs", price, "TK")
	}
}
