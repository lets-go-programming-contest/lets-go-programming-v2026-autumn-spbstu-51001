package main

import (
	"errors"
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

var (
	errEmplCount = errors.New("employees count must be between 1 and 1000")
	errOffCount  = errors.New("number of offices must be between 1 and 1000")
	errOperator  = errors.New("operator must be <= or >=")
)

func tempCalc(operator string, temp int, minTemp *int, maxTemp *int) int {
	if operator == ">=" {
		if temp > *minTemp {
			*minTemp = temp
		}
	} else if operator == "<=" {
		if temp < *maxTemp {
			*maxTemp = temp
		}
	}

	if *minTemp > *maxTemp {
		return -1
	}

	return *minTemp
}

func main() {
	var offCount uint

	if _, err := fmt.Scan(&offCount); err != nil || offCount == 0 || offCount > 1000 {
		fmt.Println("Invalid argument:", errOffCount)

		return
	}

	for range offCount {
		var emplCount uint

		if _, err := fmt.Scan(&emplCount); err != nil || emplCount == 0 || emplCount > 1000 {
			fmt.Println("Invalid argument:", errEmplCount)

			return
		}

		minTemp, maxTemp := defaultMinTemp, defaultMaxTemp

		for range emplCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("input error:", err)

				return
			}

			if operator != "<=" && operator != ">=" {
				fmt.Println("operator error:", errOperator)

				return
			}

			fmt.Println(tempCalc(operator, temp, &minTemp, &maxTemp))
		}
	}
}
