func maxSubarrayLength(nums []int, k int) int {
    maxLen, curLen := 0, 0
    left := 0
    counts := make(map[int]int)
    for right := 0; right < len(nums); right ++ {
        counts[nums[right]] += 1
        for counts[nums[right]] > k {
            counts[nums[left]] -= 1
            left += 1
        }
        curLen = right - left + 1
        if curLen > maxLen {
            maxLen = curLen
        }
    }
    return maxLen
}