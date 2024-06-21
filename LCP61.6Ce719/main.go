package main

import "fmt"

func main() {
	fmt.Println(temperatureTrend([]int{21, 18, 18, 18, 31}, []int{34, 32, 16, 16, 17}))
}

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	res := 0
	start := 0
	for i := 1; i < len(temperatureA); i++ {
		if (temperatureA[i]-temperatureA[i-1])*(temperatureB[i]-temperatureB[i-1]) > 0 || ((temperatureA[i]-temperatureA[i-1]) == 0 && (temperatureB[i]-temperatureB[i-1]) == 0) {
			res = max(res, i-start)
		} else {
			start = i
		}
	}
	return res
}
