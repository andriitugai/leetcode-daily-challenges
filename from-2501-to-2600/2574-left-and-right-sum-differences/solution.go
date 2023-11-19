func leftRightDifference(nums []int) []int {
    total := 0
    for _, num := range nums {
        total += num
    }
    result := make([]int, len(nums))
    left := 0
    right := total - nums[0]
    result[0] = total - nums[0]

    for i := 1; i < len(nums); i ++ {
        left += nums[i-1]
        right -= nums[i]
        result[i] = left - right
        if result[i] < 0 {
            result[i] = -result[i]
        }
    }
    return result
}