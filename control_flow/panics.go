package control_flow

import "fmt"

func Panic() {
	dividend, divisor := 10, 0
	result := dividend / divisor
	fmt.Println("This will cause a panic.")
	fmt.Println("And, we will never see this result:", result)
}

func PanicAndRecover() {
	var dividend, divisor float32 = 10.0, 0.0
	result := divide(dividend, divisor)
	fmt.Printf("%v / %v = %v\n", dividend, divisor, result)
}

func divide(dividend, divisor float32) float32 {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()

	return dividend / divisor
}
