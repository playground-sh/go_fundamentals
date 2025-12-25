package code_organization

import "fmt"

// dataStore uses K (comparable) for the ID and V (any) for the actual data.
type dataStore[K comparable, V any] struct {
	data map[K]V
}

func newDataStore[K comparable, V any]() *dataStore[K, V] {
	return &dataStore[K, V]{
		data: make(map[K]V),
	}
}

// Save takes a key and any value
func (ds *dataStore[K, V]) save(id K, item V) {
	ds.data[id] = item
}

// get retrieves the value by key
func (ds *dataStore[K, V]) get(id K) (V, bool) {
	val, exists := ds.data[id]
	return val, exists
}

type student struct {
	name string
}

func GenericFunctionWithConstraints() {
	// Example 1: Keys are ints, Values are Strings
	nameCache := newDataStore[int, string]()
	nameCache.save(101, "Alice")
	name, _ := nameCache.get(101)
	fmt.Println("Student Name:", name)

	// Example 2: Keys are strings, Values are Structs (any)
	studentCache := newDataStore[string, student]()
	studentCache.save("Albert", student{name: "Einstein"})
	stu, _ := studentCache.get("Albert")
	fmt.Println("Student Name:", stu.name)
}
