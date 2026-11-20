package main

import (
	"errors"
	"fmt"
)

var (
	errEmplCount = errors.New("employees  must be != 0 and <= 1000")
	errOffCount  = errors.New("number of offices must be != 0 and <= 1000")
	errOp        = errors.New("operator must be <= or >=")
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

func TempCalc(op string, temp int, minTemp *int, maxTemp *int) int {
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

	if _, err := fmt.Scan(&offCount); err != nil {
		fmt.Println(err)

		return
	}

	if offCount == 0 || offCount > 1000 {
		fmt.Println("offCount err: ", errOffCount)

		return
	}

	var i uint
	for ; i < offCount; i++ {
		var emplCount uint

		if _, err := fmt.Scan(&emplCount); err != nil {
			fmt.Println(err)

			return
		}

		if emplCount == 0 || emplCount > 1000 {
			fmt.Println("EmplCount err: ", errEmplCount)

			return
		}

		minTemp, maxTemp := defaultMinTemp, defaultMaxTemp

		var j uint
		for ; j < emplCount; j++ {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println(err)

				return
			}

			if operator != "<=" && operator != ">=" {
				fmt.Println("op err: ", errOp)

				return
			}

			fmt.Println(TempCalc(operator, temp, &minTemp, &maxTemp))
		}
	}
}
