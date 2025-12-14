package control_flow

import "fmt"

func SwitchStatements() {
	var choice string
	fmt.Print("What would you like me to do? ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println(err)
	}

	switch choice {
	case "greet":
		fmt.Println("Hello Dear!")
	case "yell":
		fmt.Println("DON'T RUN")
	case "sing":
		fmt.Println("La la la la la")
	default:
		fmt.Println("You doing alright?")
	}
}
