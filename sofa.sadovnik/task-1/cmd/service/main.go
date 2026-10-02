package main

import "fmt"

func main() {
	var frstArg, scndArg int
	_, err := fmt.Scan(&frstArg)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&scndArg)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operator string
	_, _ = fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println(frstArg + scndArg)
	case "-":
		fmt.Println(frstArg - scndArg)
	case "*":
		fmt.Println(frstArg * scndArg)
	case "/":
		if scndArg == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(frstArg / scndArg)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
