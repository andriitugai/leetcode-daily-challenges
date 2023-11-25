func getSumAbsoluteDifferences(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    total := 0
    for i := 0; i < n; i ++ {
        total += nums[i]
    }

    left, right := 0, total

    for i := 0; i < n; i ++ {
        result[i] = i * nums[i] - left + right - (n - i) * nums[i]
        left += nums[i]
        right -= nums[i]
    }
    return result
}