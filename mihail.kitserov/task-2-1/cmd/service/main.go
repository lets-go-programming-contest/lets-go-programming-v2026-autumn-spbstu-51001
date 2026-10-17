package main

import (
	"fmt"
)

func processRequest(operator string, num int, minNum *int, maxNum *int) int {
	if operator == ">=" {
		if num > *minNum {
			*minNum = num
		}
	} else if operator == "<=" {
		if num < *maxNum {
			*maxNum = num
		}
	}

	if *minNum > *maxNum {
		return -1
	}

	return *minNum
}

func main() {
	var officesCount int

	_, err := fmt.Scan(&officesCount)
	if err != nil {
		fmt.Println("Non-correct count offices")

		return
	}

	i := 0
	for ; i < officesCount; i++ {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Non-correct count employees")

			return
		}

		minNum := 15
		maxNum := 30

		j := 0
		for ; j < employeeCount; j++ {
			var operator string

			_, err := fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Non-correct operator input")

				return
			}

			if operator != "<=" && operator != ">=" {
				fmt.Println("Non-correct operator input")

				return
			}

			var num int

			_, err = fmt.Scan(&num)
			if err != nil {
				fmt.Println("Non-correct input number")

				return
			}

			ans := processRequest(operator, num, &minNum, &maxNum)

			fmt.Println(ans)
		}
	}
}
