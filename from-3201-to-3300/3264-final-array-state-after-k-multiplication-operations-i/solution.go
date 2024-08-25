func getFinalState(nums []int, k int, multiplier int) []int {
    n := len(nums)
    findAndMul := func() {
        minIdx, minVal := 0, nums[0]
        for i := 1; i < n; i ++ {
            if nums[i] < minVal {
                minVal = nums[i]
                minIdx = i
            }
        }
        nums[minIdx] *= multiplier
    }
    for i := 0; i < k; i ++ {
        findAndMul()
    }
    return nums
}