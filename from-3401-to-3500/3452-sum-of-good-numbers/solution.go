func sumOfGoodNumbers(nums []int, k int) int {
    result := 0
    n := len(nums)
    for i := 0; i < n; i ++ {
        if i - k >= 0 && nums[i - k] >= nums[i] || i + k < n && nums[i + k] >= nums[i] {
            continue
        }
        result += nums[i]
    }
    return result
}