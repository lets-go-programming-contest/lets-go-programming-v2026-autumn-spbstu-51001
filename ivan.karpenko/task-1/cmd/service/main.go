package main

import "fmt"

func main() {
	var num1 int
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var num2 int
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var op string
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	var res int
	switch op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = num1 / num2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(res)
}
