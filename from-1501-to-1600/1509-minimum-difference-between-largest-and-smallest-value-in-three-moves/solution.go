func minDifference(nums []int) int {
    n := len(nums)
    if n <= 4 {
        return 0
    }
    sort.Ints(nums)
    result := math.MaxInt
    for left := 0; left < 4; left ++ {
        diff := nums[n - 4 + left] - nums[left]
        if diff < result {
            result = diff
        }
    }
    return result
}