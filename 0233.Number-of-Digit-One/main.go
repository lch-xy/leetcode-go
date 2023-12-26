package main

func main() {
	println(countDigitOne(13))
}
func countDigitOne(n int) int {
	high := n
	low := 0
	cur := 0
	count := 0
	mult := 1
	for high != 0 || cur != 0 {
		cur = high % 10
		high /= 10
		if cur == 0 {
			// 如果是0，那么从0开始到高位，一定有high * mult个1
			// 例如 100，我们在计算第二位是=时，虽然是0，需要从0走到9后才会往前进1
			// 所以肯定有1个1，这个1是前面的1带给他的
			count += high * mult
		} else if cur == 1 {
			// high*mult是当前的倍数，所以肯定是包含这个1的
			// 而low代表是低位，只要low有多少个，前面那位的1就有多少个，因为low没走一步都会带上他
			// 为什么要+1？因为当前是1，单独算
			count += high*mult + 1 + low
		} else {
			// 为什么是high+1？
			// 因为对于当前这一位来说，如果大于1，那么肯定就包含了1的一切
			// 就含义上来说和进一位没多大区别，因为从2-9是没啥用的，有用的只有1
			// 为什么cur=1时是要分开加？
			// 因为并不包含1的全部，例如low=1时，其实只包含的一点点
			count += (high + 1) * mult
		}
		low += cur * mult
		mult *= 10
	}
	return count
}
