package main

import "fmt"

func main() {
	type TestCase struct {
		mainTank       int
		additionalTank int
		output         int
	}

	testCases := []TestCase{
		{mainTank: 1, additionalTank: 5, output: 10},
		{mainTank: 9, additionalTank: 2, output: 110},
	}

	for i, tc := range testCases {
		result := distanceTraveled(tc.mainTank, tc.additionalTank)
		if result == tc.output {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed: expected %d, got %d\n", i+1, tc.output, result)
			return
		}
	}
}

// simulation step
func distanceTraveled(mainTank int, additionalTank int) int {
	distance := 0
	for mainTank >= 5 {
		mainTank -= 5
		distance += 5
		if additionalTank >= 1 {
			mainTank += 1
			additionalTank -= 1
		}
	}
	return (distance + mainTank) * 10
}
