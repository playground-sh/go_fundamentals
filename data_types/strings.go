package data_types

import (
	"fmt"
	"strings"
)

func StringData() {
	var name string
	name = "Thor Odinson"
	fmt.Println("The name is", name)
	fmt.Printf("Name is %d characters long\n", len(name))

	fullName := strings.Split(name, " ")
	firstName, lastName := fullName[0], fullName[1]
	fmt.Printf("First name is `%s` and last name is `%s`\n", firstName, lastName)
}
