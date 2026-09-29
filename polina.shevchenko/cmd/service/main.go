package main

import "fmt"

func main() {
	var first, second int
	var operation string

	switch operation {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		fmt.Println(first / second)
	}
}
