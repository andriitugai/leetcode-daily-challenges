func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    n := len(nums)
    countLessThan := func(limit int) int64 {
        var result int64
        left, right := 0, n - 1
        for left < right {
            for nums[left] + nums[right] > limit && left < right {
                right -= 1
            }
            result += int64(right - left)
            left += 1
        }
        return result
    }
    return countLessThan(upper) - countLessThan(lower - 1)
}