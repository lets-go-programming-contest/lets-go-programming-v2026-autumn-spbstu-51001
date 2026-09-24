package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1, input2, operator string
	if _, err := fmt.Scan(&input1, &input2, &operator); err != nil {
		return
	}

	number1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	number2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(number1 / number2)
	default:
		fmt.Println("Invalid operation")
	}
}
