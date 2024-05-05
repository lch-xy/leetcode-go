package main

func decrypt(code []int, k int) []int {
	length := len(code)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		if k == 0 {
			res[i] = 0
		} else if k > 0 {
			sum := 0
			for j := 1; j <= k; j++ {
				sum += code[(i+j+length)%length]
			}
			res[i] = sum
		} else if k < 0 {
			sum := 0
			for j := 1; j <= -k; j++ {
				sum += code[(i-j+length)%length]
			}
			res[i] = sum
		}
	}
	return res
}
