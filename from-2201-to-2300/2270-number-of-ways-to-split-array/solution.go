func waysToSplitArray(nums []int) int {
    leftSum, rightSum := 0, 0
    for _, num := range nums {
        rightSum += num
    }
    result := 0
    for i := 0; i < len(nums) - 1; i ++ {
        leftSum += nums[i]
        rightSum -= nums[i]
        if leftSum >= rightSum {
            result += 1
        }
    }
    return result
}