package code_organization

import "fmt"

type info interface {
	showInfo()
}

type Student struct {
	name string
	cgpa float32
}

func (s *Student) showInfo() {
	fmt.Println("Student Info")
	fmt.Println("Name:", s.name)
	fmt.Println("CGPA:", s.cgpa)
}

type Engineer struct {
	name       string
	department string
}

func (e *Engineer) showInfo() {
	fmt.Println("Engineer Info")
	fmt.Println("Name:", e.name)
	fmt.Println("Department:", e.department)
}

func InterfaceDemo() {
	s1 := Student{
		name: "Larry",
		cgpa: 3.50,
	}

	e1 := Engineer{
		name:       "Sam",
		department: "CSE",
	}

	e2 := Engineer{
		name:       "Bran",
		department: "EEE",
	}

	// add different concrete types to a list of the generic type (`info`)
	// var people []info
	// people = append(people, &s1)
	// people = append(people, &e1)
	// people = append(people, &e2)

	// alternative syntax for the above
	people := []info{&s1, &e1, &e2}

	for idx, person := range people {
		if idx == 0 {
			fmt.Println("--------------------------------")
		}
		person.showInfo()
		fmt.Println("--------------------------------")
	}
}
