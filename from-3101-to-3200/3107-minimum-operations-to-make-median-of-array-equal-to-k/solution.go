func minOperationsToMakeMedianK(nums []int, k int) int64 {
    var result int64 = 0
    sort.Ints(nums)
    n := len(nums)
    medIdx := n / 2

    if nums[medIdx] > k {
        for i := medIdx; i >= 0; i -- {
            if nums[i] > k {
                result += int64(nums[i] - k)
            }
        }
    } else if nums[medIdx] < k {
        for i := medIdx; i < n; i ++ {
            if nums[i] < k {
                result += int64(k - nums[i])
            }
        }
    }
    return result
}