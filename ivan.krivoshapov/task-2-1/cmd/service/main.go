package main

import "fmt"

const (
	lowestTemp  = 15
	highestTemp = 30
)

func pickTemperature(sign string, value int, lower, upper *int) int {
	switch sign {
	case ">=":
		if value > *lower {
			*lower = value
		}
	case "<=":
		if value < *upper {
			*upper = value
		}
	}

	if *lower > *upper {
		return -1
	}

	return *lower
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid departments count")

		return
	}

	for range departments {
		var staff int

		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Println("Invalid employees count")

			return
		}

		lower, upper := lowestTemp, highestTemp

		for range staff {
			var sign string

			var value int

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("Invalid input")

				return
			}

			if sign != ">=" && sign != "<=" {
				fmt.Println("Invalid operator")

				return
			}

			fmt.Println(pickTemperature(sign, value, &lower, &upper))
		}
	}
}
