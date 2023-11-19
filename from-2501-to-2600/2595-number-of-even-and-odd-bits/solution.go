func evenOddBit(n int) []int {
	ans := []int{0, 0}
	for i := 32; i >= 0; i-- {
		if val := n & (1 << i); val > 0 {
			if i%2 == 0 {
				ans[0]++
			} else {
				ans[1]++
			}
		}
	}
	return ans
}