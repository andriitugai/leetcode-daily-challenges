func minimumSubarrayLength(nums []int, k int) int {
    minLen := 100
    for i := 0; i < len(nums); i ++ {
        or := 0
        for j := i; j < len(nums); j ++ {
            or |= nums[j]
            if or >= k && j - i + 1 < minLen {
                minLen = j - i + 1
            }
        }
    }
    if minLen == 100 {
        return -1
    }
    return minLen
}