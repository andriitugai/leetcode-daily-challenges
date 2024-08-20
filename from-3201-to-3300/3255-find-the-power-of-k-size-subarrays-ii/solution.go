func resultsArray(nums []int, k int) []int {
    n := len(nums)
    result := make([]int, n - k + 1)
    score := 0
    for i := 0; i < n; i ++ {
        if i > 0 && nums[i - 1] + 1 == nums[i] {
            score += 1
        } else {
            score = 0
        }
        if i >= k - 1 {
            if score >= k - 1 {
                result[i - k + 1] = nums[i]
            } else {
                result[i - k + 1] = -1
            }
        }
    }
    return result
}