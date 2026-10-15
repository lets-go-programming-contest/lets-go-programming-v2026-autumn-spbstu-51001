package main

import "fmt"

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		minTemperature := 15
		maxTemperature := 30

		for range employeesCount {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				return
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
	}
}
