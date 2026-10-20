package main

import "fmt"

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid departments count")
		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employees count")
			return
		}
		if !processDepartment(employeesCount) {
			return
		}
	}
}

func processDepartment(employeesCount int) bool {
	minTemperature := 15
	maxTemperature := 30

	for range employeesCount {
		var (
			operation   string
			temperature int
		)

		if _, err := fmt.Scan(&operation); err != nil {
			fmt.Println("Invalid operation input")
			return false
		}
		if _, err := fmt.Scan(&temperature); err != nil {
			fmt.Println("Invalid temperature requirement")
			return false
		}

		if operation == ">=" && temperature > minTemperature {
			minTemperature = temperature
		} else if operation == "<=" && temperature < maxTemperature {
			maxTemperature = temperature
		}

		if minTemperature > maxTemperature {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemperature)
		}
	}

	return true
}
