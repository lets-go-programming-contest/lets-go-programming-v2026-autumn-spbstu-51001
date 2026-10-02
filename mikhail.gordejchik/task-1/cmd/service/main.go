package main

import (
	"fmt"
)

func main() {
	var firstVariable int
	_, err := fmt.Scan(&firstVariable)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var secondVariable int
	_, err = fmt.Scan(&secondVariable)
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
		fmt.Println(firstVariable + secondVariable)
	case "-":
		fmt.Println(firstVariable - secondVariable)
	case "*":
		fmt.Println(firstVariable * secondVariable)
	case "/":
		if secondVariable == 0 {
			fmt.Println("Division by zero")
			break
		}
		fmt.Println(firstVariable / secondVariable)
	default:
		fmt.Println("Invalid operation")
	}
}
