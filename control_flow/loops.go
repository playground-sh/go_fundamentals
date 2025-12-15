package control_flow

import "fmt"

func InfiniteLoop() {
	fmt.Println("Infinite Loop")
	i := 0
	for {
		i++
		fmt.Println(i)

		// safety breaks
		if i == 5 {
			break
		}
	}
}

func LoopUntil() {
	fmt.Println("Loop Until")
	i := 6
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func CounterBasedLoop() {
	fmt.Println("Counter Based Loop (a.k.a. Traditional For Loop)")
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func RangeBasedLoop() {
	fmt.Println("Range Based Loop")
	for i := range 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func LoopingOverCollections() {
	fmt.Println("Looping over a collection")

	type student struct {
		name       string
		department string
		cgpa       float32
	}

	var students = []student{
		{name: "Sofia", department: "CSE", cgpa: 4.0},
		{name: "Miller", department: "BBA", cgpa: 3.7},
		{name: "Scott", department: "EEE", cgpa: 3.97},
	}

	for idx, member := range students {
		fmt.Printf("ID: %d, Name: %s, Department: %s, CGPA: %.2f\n",
			idx+1, member.name, member.department, member.cgpa)
	}
}
