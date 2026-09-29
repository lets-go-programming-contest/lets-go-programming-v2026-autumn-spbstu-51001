package main

import "fmt"

func main() {
	var first, second int
	var operation string

	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
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
		if second == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(first / second)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
