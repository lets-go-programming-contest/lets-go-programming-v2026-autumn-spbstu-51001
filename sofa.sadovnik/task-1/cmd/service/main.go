package main

import "fmt"

func main() {
	var frstArg, scndArg int

	var operator string

	switch operator {
	case "+":
		fmt.Println(frstArg + scndArg)
	case "-":
		fmt.Println(frstArg - scndArg)
	case "*":
		fmt.Println(frstArg * scndArg)
	case "/":
		fmt.Println(frstArg / scndArg)
	}
}
