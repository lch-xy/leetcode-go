package main

import "math"

func average(salary []int) float64 {
	maxSalary := math.MinInt32
	minSalary := math.MaxInt32

	total := 0
	for i := 0; i < len(salary); i++ {
		if maxSalary < salary[i] {
			maxSalary = salary[i]
		}
		if minSalary > salary[i] {
			minSalary = salary[i]
		}
		total += salary[i]
	}

	return float64(total-maxSalary-minSalary) / float64((len(salary) - 2))
}
