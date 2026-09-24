package main

import "fmt"

func main() {
	var firstArg int64
	_, err := fmt.Scanln(&firstArg)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var secondArg int64
	_, err = fmt.Scanln(&secondArg)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var operator string
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(firstArg + secondArg)
	case "-":
		fmt.Println(firstArg - secondArg)
	case "*":
		fmt.Println(firstArg * secondArg)
	case "/":
		if secondArg == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(firstArg / secondArg)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
