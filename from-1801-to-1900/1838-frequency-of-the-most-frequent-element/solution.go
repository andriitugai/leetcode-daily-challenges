func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)

    left := 0
    result, sum := 1, 0

    for right := 0; right < len(nums); right ++ {
        sum += nums[right]
        if nums[right] * (right - left + 1) > sum + k {
            sum -= nums[left]
            left += 1
        }
        if right - left + 1 > result {
            result = right - left + 1
        }
    }
    return result
}