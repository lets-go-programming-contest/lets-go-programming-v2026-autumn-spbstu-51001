package main

import (
	"errors"
	"fmt"

	"github.com/CWSHBR/task-2-1/internal/temperature"
)

const (
	minCount = 1
	maxCount = 1000
)

var errCountOutOfRange = errors.New("value is out of range")

func readCount(name string) (int, error) {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("read %s: %w", name, err)
	}

	if count < minCount || count > maxCount {
		return 0, fmt.Errorf("%s %d: %w [%d, %d]", name, count, errCountOutOfRange, minCount, maxCount)
	}

	return count, nil
}

func processDepartment() error {
	employees, err := readCount("employees count")
	if err != nil {
		return err
	}

	tempRange := temperature.NewRange()

	for range employees {
		var (
			operator string
			value    int
		)

		_, err = fmt.Scan(&operator, &value)
		if err != nil {
			return fmt.Errorf("read temperature request: %w", err)
		}

		tempRange, err = tempRange.Apply(operator, value)
		if err != nil {
			return fmt.Errorf("apply temperature request: %w", err)
		}

		fmt.Println(tempRange.Optimal())
	}

	return nil
}

func main() {
	departments, err := readCount("departments count")
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	for range departments {
		err = processDepartment()
		if err != nil {
			fmt.Println("Invalid input:", err)

			return
		}
	}
}
