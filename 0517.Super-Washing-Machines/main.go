package main

import "fmt"

func main() {
	testCases := []struct {
		name     string
		machines []int
		expected int
	}{
		{"All machines have the same amount", []int{4, 4, 4, 4}, 0},
		{"One machine has all the clothes", []int{0, 0, 0, 12}, 9},
		{"Example case", []int{0, 0, 11, 5}, 8},
		{"Example case", []int{0, 3, 0}, 2},
	}

	for _, tc := range testCases {
		got := findMinMoves(tc.machines)
		if got != tc.expected {
			fmt.Printf("findMinMoves(%v) = %d; want %d。", tc.machines, got, tc.expected)
		}
	}
}

// 思路：[0, 0, 11, 5] -> [4, 4, 4, 4]
// 那么我们将二者做差，得到[-4, -4, 7, 1]，这里负数表示当前洗衣机还需要的衣服数，正数表示当前洗衣机多余的衣服数。
// 我们要做的是要将这个差值数组每一项都变为0，对于第一个洗衣机来说，需要四件衣服可以从第二个洗衣机获得，那么就可以把-4移给二号洗衣机，那么差值数组变为[0, -8, 7, 1]，此时二号洗衣机需要八件衣服，那么至少需要移动8次。
// 然后二号洗衣机把这八件衣服从三号洗衣机处获得，那么差值数组变为[0, 0, -1, 1]，此时三号洗衣机还缺1件，就从四号洗衣机处获得，此时差值数组成功变为了[0, 0, 0, 0]，成功。
// 那么移动的最大次数就是差值数组中出现的绝对值最大的数字，8次
func findMinMoves(machines []int) int {
	len := len(machines)
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%len != 0 {
		return -1
	}
	res, diff, target := 0, 0, sum/len
	for i := 0; i < len; i++ {
		diff = -(machines[i] - target - diff)
		// 因为每次只能转移1件，所以要和machines[i]-target进行比较，不然无法通过case[0,3,0]
		res = max(res, machines[i]-target, diff, -diff)
	}
	return res
}
