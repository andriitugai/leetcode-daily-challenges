func minIncrementForUnique(nums []int) int {
    sort.Ints(nums)
    var delta int
    result := 0
    for i := 1; i < len(nums); i ++ {
        if nums[i] <= nums[i - 1] {
            delta = nums[i - 1] - nums[i] + 1
            result += delta
            nums[i] += delta
        }
    }
    return result
}