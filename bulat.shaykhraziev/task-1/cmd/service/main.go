package main

import "fmt"

func main() {
	var first int

	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")

		return
	}

	var second int

	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")

		return
	}

	var operation string

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")

		return
	}

	switch operation {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")

			return
		}

		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
