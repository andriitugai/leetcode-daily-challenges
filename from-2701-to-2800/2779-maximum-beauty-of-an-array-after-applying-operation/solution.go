func maximumBeauty(nums []int, k int) int {
    sort.Ints(nums)
    result := 0
    left := 0
    for right := 0; right < len(nums); right ++ {
        for nums[right] - nums[left] > 2 * k {
            left += 1
        }
        curLen := right - left + 1
        if curLen > result {
            result = curLen
        }
    }
    return result
}