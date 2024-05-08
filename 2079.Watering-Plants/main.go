package main

func wateringPlants(plants []int, capacity int) int {
	total := 0
	curCapacity := capacity
	for i := 0; i < len(plants); i++ {
		total++
		if curCapacity >= plants[i] {
			curCapacity -= plants[i]
		} else {
			total += 2 * i
			total--
			i--
			curCapacity = capacity
		}
	}
	return total
}
