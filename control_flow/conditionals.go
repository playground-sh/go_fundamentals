package control_flow

import "fmt"

func IfElse() {
	fmt.Println("If Statements")

	var guess int
	fmt.Println("Guess the magic number:")
	_, err := fmt.Scanln(&guess)
	if err != nil {
		fmt.Println(err)
	}

	if magicNumber := 200; guess < magicNumber {
		fmt.Println("You guessed less!")
	} else if guess > magicNumber {
		fmt.Println("You guessed too high!")
	} else {
		fmt.Println("Congrats!!! You are correct, genious!")
	}
}
