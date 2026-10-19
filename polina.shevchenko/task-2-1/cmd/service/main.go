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
	errOp        = errors.New("operator must be <= or >=")
)

func tempCalc(op string, temp int, minTemp *int, maxTemp *int) int {
	if op == ">=" {
		if temp > *minTemp {
			*minTemp = temp
		}
	} else if op == "<=" {
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
		fmt.Println("offCount err:", errOffCount)
		return
	}

	for i := uint(0); i < offCount; i++ {
		var emplCount uint

		if _, err := fmt.Scan(&emplCount); err != nil || emplCount == 0 || emplCount > 1000 {
			fmt.Println("emplCount err:", errEmplCount)
			return
		}

		minTemp, maxTemp := defaultMinTemp, defaultMaxTemp
		for j := uint(0); j < emplCount; j++ {
			var op string
			var temp int

			if _, err := fmt.Scan(&op, &temp); err != nil {
				fmt.Println("input err:", err)
				return
			}

			if op != "<=" && op != ">=" {
				fmt.Println("operator err:", errOp)
				return
			}

			fmt.Println(tempCalc(op, temp, &minTemp, &maxTemp))
		}
	}
}
