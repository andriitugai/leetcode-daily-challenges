func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	isArithmetic := func(left, right int) bool {
		length := right - left + 1
		if length < 3 {
			return true
		}
		maxVal, minVal := nums[left], nums[left]
		for i := left + 1; i <= right; i++ {
			if nums[i] > maxVal {
				maxVal = nums[i]
			} else if nums[i] < minVal {
				minVal = nums[i]
			}
		}

		diff := maxVal - minVal
		if diff == 0 {
			return true
		}
		if diff%(length-1) != 0 {
			return false
		}

		delta := diff / (length - 1)
		visited := make(map[int]bool)
		for i := left; i <= right; i++ {
			if (nums[i]-minVal)%delta != 0 {
				return false
			}
			idx := (nums[i] - minVal) / delta
			if visited[idx] {
				return false
			}
			visited[idx] = true
		}
		return true
	}

	result := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		result[i] = isArithmetic(l[i], r[i])
	}
	return result
}
