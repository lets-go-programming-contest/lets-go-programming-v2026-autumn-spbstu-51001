package main

import "fmt"

const (
	minTempInit = 15
	maxTempInit = 30
)

func calcTemp(operation string, temp, minTemp, maxTemp int) (int, int) {
	if operation == "<=" && temp < maxTemp {
		maxTemp = temp
	}

	if operation == ">=" && temp > minTemp {
		minTemp = temp
	}

	return minTemp, maxTemp
}

func main() {
	var amountOffice int

	_, err := fmt.Scan(&amountOffice)
	if err != nil {
		fmt.Println("Error: failed to read number of offices")

		return
	}

	for range amountOffice {
		var amountEmployee int

		_, err = fmt.Scan(&amountEmployee)
		if err != nil {
			fmt.Println("Error: failed to read number of employees")

			return
		}

		minTemp := minTempInit
		maxTemp := maxTempInit

		for range amountEmployee {
			var operation string

			_, err = fmt.Scan(&operation)
			if err != nil {
				fmt.Println("Error: failed to read operation")

				return
			}

			if operation != "<=" && operation != ">=" {
				fmt.Println("Error: invalid operation, expected '<=' or '>='")

				return
			}

			var num int

			_, err = fmt.Scan(&num)
			if err != nil {
				fmt.Println("Error: failed to read temperature")

				return
			}

			minTemp, maxTemp = calcTemp(operation, num, minTemp, maxTemp)

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
