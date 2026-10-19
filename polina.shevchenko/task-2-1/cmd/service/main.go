package main

import "fmt"

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

func tempCalc() {}
func main() {
	var offCount uint

	if _, err := fmt.Scan(&offCount); err != nil || offCount == 0 || offCount > 1000 {
		fmt.Println()
		return
	}

	for i := uint(0); i < offCount; i++ {
		var emplCount uint

		if _, err := fmt.Scan(&emplCount); err != nil || emplCount == 0 || emplCount > 1000 {
			fmt.Println()
			return
		}

		for j := uint(0); j < emplCount; j++ {
			var op string
			var temp int

			if _, err := fmt.Scan(&op, &temp); err != nil {
				fmt.Println()
				return
			}

			if op != "<=" && op != ">=" {
				fmt.Println()
				return
			}

			fmt.Println(tempCalc())
		}
	}
}
