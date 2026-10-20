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

		switch operation {
		case ">=":
			if temperature > minTemperature {
				minTemperature = temperature
			}
		case "<=":
			if temperature < maxTemperature {
				maxTemperature = temperature
			}
		default:
			fmt.Println("Invalid operation")
			return false
		}

		if minTemperature > maxTemperature {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemperature)
		}
	}

	return true
}
