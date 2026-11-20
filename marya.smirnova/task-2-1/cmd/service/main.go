package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func processDepartment(employeeCount int) bool {
	lower := minTemp
	upper := maxTemp

	for range employeeCount {
		var (
			sign  string
			value int
		)

		if _, err := fmt.Scan(&sign, &value); err != nil {
			fmt.Println("Invalid employee input")

			return false
		}

		switch sign {
		case ">=":
			if value > lower {
				lower = value
			}

		case "<=":
			if value < upper {
				upper = value
			}

		default:
			fmt.Println("Invalid operation")

			return false
		}

		if lower > upper {
			fmt.Println(-1)
		} else {
			fmt.Println(lower)
		}
	}

	return true
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Invalid departments count")

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Invalid employees count")

			return
		}

		if !processDepartment(employeeCount) {
			return
		}
	}
}
