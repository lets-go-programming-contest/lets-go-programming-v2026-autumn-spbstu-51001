package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30

	noSuitableTemp = -1
)

var errInvalidOperator = errors.New("operator must be >= or <=")

func readCount(name string) (int, error) {
	var count int

	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("reading %s: %w", name, err)
	}

	return count, nil
}

func applyWish(operator string, temp, low, high int) (int, int, error) {
	switch operator {
	case ">=":
		low = max(low, temp)
	case "<=":
		high = min(high, temp)
	default:
		return low, high, fmt.Errorf("invalid operator %q: %w", operator, errInvalidOperator)
	}

	return low, high, nil
}

func optimalTemp(low, high int) int {
	if low > high {
		return noSuitableTemp
	}

	return low
}

func processDepartment() error {
	employees, err := readCount("employees count")
	if err != nil {
		return err
	}

	low, high := minTemp, maxTemp

	for range employees {
		var (
			operator string
			temp     int
		)

		if _, err := fmt.Scan(&operator, &temp); err != nil {
			return fmt.Errorf("reading employee request: %w", err)
		}

		low, high, err = applyWish(operator, temp, low, high)
		if err != nil {
			return err
		}

		fmt.Println(optimalTemp(low, high))
	}

	return nil
}

func main() {
	departments, err := readCount("departments count")
	if err != nil {
		fmt.Println("Error while reading input:", err)

		return
	}

	for range departments {
		if err := processDepartment(); err != nil {
			fmt.Println("Error while reading input:", err)

			return
		}
	}
}
