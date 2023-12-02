func findKOr(nums []int, k int) int {
    result := 0
    mask := 1
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, x := range nums {
			if x & mask != 0 {
				cnt++
			}
		}
		if cnt >= k {
			result += mask
		}
		mask = mask << 1
	}
	return result
}