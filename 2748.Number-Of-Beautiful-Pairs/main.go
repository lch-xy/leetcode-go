package main

func main() {
	println(countBeautifulPairs([]int{756, 1324, 2419, 495, 106, 111, 1649, 1474, 2001, 1633, 273, 1804, 2102, 1782, 705, 1529, 1761, 1613, 111, 186, 412}))
}

func countBeautifulPairs(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(getHead(nums[i]), nums[j]%10) == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func getHead(a int) int {
	for a >= 10 {
		a = a / 10
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
