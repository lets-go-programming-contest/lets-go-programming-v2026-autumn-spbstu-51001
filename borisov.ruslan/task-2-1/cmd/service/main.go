package main

import "fmt"

func changeClimatHandler(operator string, minTemp, maxTemp *int, num int) int {
	switch operator {
	case ">=":
		if num > *minTemp {
			*minTemp = num
		}

	case "<=":
		if num < *maxTemp {
			*maxTemp = num
		}
	}

	if *minTemp > *maxTemp {
		return -1
	}

	return *minTemp
}

func main() {
	var countDepartments int

	if _, err := fmt.Scan(&countDepartments); err != nil {
		fmt.Println("Invalid count of departments")

		return
	}

	for range countDepartments {
		var countEmployees int

		if _, err := fmt.Scan(&countEmployees); err != nil {
			fmt.Println("Invalid count of employees")

			return
		}

		minTemp := 15
		maxTemp := 30

		for range countEmployees {
			var operator string

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator")

				return
			}

			if operator != ">=" && operator != "<=" {
				fmt.Println("Invalid operator")

				return
			}

			var num int

			if _, err := fmt.Scan(&num); err != nil {
				fmt.Println("Invalid number")

				return
			}

			corrTemp := changeClimatHandler(
				operator, &minTemp, &maxTemp, num,
			)

			fmt.Println(corrTemp)
		}
	}
}
