package data_types

import "fmt"

type student struct {
	name       string
	department string
	gpa        float32
}

func Structs() {
	amy := student{
		"Amy Weinberg",
		"CSE",
		3.50,
	}
	fmt.Println(amy)

	alex := student{
		name:       "Alex Miller",
		department: "BBA",
		gpa:        3.81,
	}
	fmt.Println(alex)

	// access members individually
	fmt.Println(alex.name)
	fmt.Println(alex.department)
	fmt.Println(alex.gpa)
}
