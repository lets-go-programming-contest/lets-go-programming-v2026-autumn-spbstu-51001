package main

import "fmt"

func main() {
	var first, second int
	var operation string

	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
	}

	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
	}

	_, _ = fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
