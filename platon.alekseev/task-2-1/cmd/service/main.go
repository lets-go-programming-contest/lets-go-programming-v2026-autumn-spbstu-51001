package main

import "fmt"

func main() {
	var departmentsCount int
	fmt.Scan(&departmentsCount)

	for department := 0; department < departmentsCount; department++ {
		var employeesCount int
		fmt.Scan(&employeesCount)

		for employee := 0; employee < employeesCount; employee++ {
			var operation string
			var temperature int
			fmt.Scan(&operation, &temperature)
		}
	}
}
