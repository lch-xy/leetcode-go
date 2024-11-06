package main

func losingPlayer(x int, y int) string {
	locaion := 0
	names := []string{"Alice", "Bob"}
	if !(x >= 1 && y >= 4) {
		return names[locaion+1]
	}
	canFind := true
	for canFind {
		locaion++
		if !(x >= 1 && y >= 4) {
			canFind = false
			break
		}
		x -= 1
		y -= 4
	}
	return names[locaion]
}

func losingPlayer2(x int, y int) string {
	names := []string{"Bob", "Alice"}
	cnt := min(x, y/4)
	return names[cnt%2]
}
